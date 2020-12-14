package src

import (
	"encoding/json"
	"strconv"
	"strings"
	"time"
)

// Загрузить кэш запросов
func LoadCache() {
	if j := A.Preferences().StringWithFallback("cache_requests", ""); j != "" {
		// загрузить весь кэш, отобрать только актуальный (за сегодня) и пересохранить
		existsRoutes := map[string][]route{}
		err := json.Unmarshal([]byte(j),&existsRoutes)
		for k,v := range existsRoutes {
			if strings.Contains(k,time.Now().Format("2006-01-02")) {
				routesCache[k] = v
			}
		}
		IsErr(err)
		saveCache()
	}
}

// Сохранить кэш запросов
func saveCache() {
	data,err := json.Marshal(routesCache)
	IsErr(err)
	A.Preferences().SetString("cache_requests", string(data))
}

// Загрузить из preferences города, маршруты между которыми запрашивались в последний раз
func LoadLastCities() (c1,c2 string) {
	s := A.Preferences().StringWithFallback("last_cities", "")
	split := strings.Split(s,"•")
	if s == "" || len(split) != 2 {
		return "Минск","Брест"
	}
	return split[0], split[1]
}

// Сохранить в preferences города, маршруты между которыми запрашивались в последний раз
func saveLastCities(c1,c2 string) {
	A.Preferences().SetString("last_cities", c1 + "•" + c2)
}

// Загрузить количество полученных диаграм
func LoadCountOfCompletedDiagrams() int {
	s := A.Preferences().StringWithFallback("count_of_completed_diagrams", "")
	r,err := strconv.Atoi(s)
	if s == "" || err != nil {
		return 0
	}
	return r
}

// Сохранить количество полученных диаграм
func saveCountOfCompletedDiagrams() {
	A.Preferences().SetString("count_of_completed_diagrams", strconv.Itoa(СколькоДиаграммПоказано))
}