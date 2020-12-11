package src

import (
	"encoding/json"
	"fmt"
	r "github.com/onix3/train-timetable/resource"
	"io/ioutil"
	"net/http"
	"sort"
	"time"
)

// Рейс состоит из названий станций и времени отправления/прибытия
type route struct{
	S1 string    `json:"s1"`
	S2 string    `json:"s2"`
	T1 time.Time `json:"t1"`
	T2 time.Time `json:"t2"`
}

// Существует ли такой запрос в кэше
func inCache(q string) bool {
	for k := range routesCache {
		if q == k {
			return true
		}
	}
	return false
}

// Получение списка рейсов в прямом и обратном направлениях,
// к тому же отсортированных по времени отправления
func getAllSortedRoutes(c1, c2, date, tt string) (allRoutes []route) {
	button1()
	routes12 := getRoutes(c1, c2, date, tt)
	button2()
	routes21 := getRoutes(c2, c1, date, tt)
	go buttonBack()
	allRoutes = append(routes12, routes21...)

	sort.Slice(allRoutes, func(i,j int) bool {
		return allRoutes[i].T1.Before(allRoutes[j].T1)
	})

	// показать в консоли
	for _,t := range allRoutes {
		fmt.Printf("%15s → %-15s %s — %s\n",
			t.S1, t.S2,
			t.T1.Format("15:04"), t.T2.Format("15:04"))
	}

	return
}

// Формирует ссылку и осуществляет запрос к API (если запроса нет в кэше)
func getRoutes(c1, c2, date, tt string) (routes []route) {
	q := fmt.Sprintf("%s%s%s%s%s%s%s",
		codeOf[c1], string(r.T), codeOf[c2], string(r.F), tt, string(r.D), date)
	url := fmt.Sprintf("%s%s",
		string(r.A), q)

	routes = getRoutesFromUrl(c1, c2,url)

	return
}

func getRoutesFromUrl(c1,c2, url string) (routes []route) {
	type Segment struct {
		Arrival   time.Time `json:"arrival"`
		Departure time.Time `json:"departure"`
	}
	var response struct {
		Segments []*Segment `json:"segments"`
	}

	j := getJson(url)

	err := json.Unmarshal(j,&response)
	IsErr(err)

	for _,s := range response.Segments {
		// почему-то Янд-кс присылает маршруты и следующих суток
		if s.Departure.Before(today0000.AddDate(0,0,1)) {
			routes = append(routes, route{
				S1: c1,
				S2: c2,
				T1: s.Departure,
				T2: s.Arrival,
			})
		}
	}

	return
}

// json по запросу
func getJson(url string) (data []byte) {
	resp,err := http.Get(url)
	IsErr(err)
	data,err = ioutil.ReadAll(resp.Body)
	IsErr(err)
	err = resp.Body.Close()
	IsErr(err)

	return
}