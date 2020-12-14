package src

import (
	"fmt"
	"fyne.io/fyne/dialog"
	"github.com/anthonynsimon/bild/transform"
	"net/http"
	"time"
)

// Проверка подключения к Интернету
func haveConnection() bool {
	client := http.Client{
		Timeout: 500*time.Millisecond,
	}
	resp,err := client.Get("https://www.google.com/")
	if err != nil {
		fmt.Println(err)
		return false
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		fmt.Println(resp.StatusCode)
		return false
	}
	return true
}

// Событие нажатия кнопки
func mainButtonClick(s1,s2,tt string) {
	q := fmt.Sprintf("%s-%s(%s)",
		s1,s2, time.Now().Format("2006-01-02"))

	var routes []route

	// кэш расписания доступен оффлайн
	// подключение в Интернету проверять, если нет в кэше
	if inCache(q) {
		button2()
		routes = routesCache[q]
		go buttonBack(time.Second)
	} else {
		PastCity1,PastCity2 = s1,s2
		saveLastCities(s1,s2)

		diagramBox.Hide()

		if s1 == s2 {
			dialog.ShowInformation("", "Города должны быть разными", W)
			return
		}

		button1()
		// проверка подключения к Интернету
		if !haveConnection() {
			dialog.ShowInformation("", "Нет подключения к Интернету", W)
			buttonBack(0)
			return
		}

		routes = getAllSortedRoutes(s1,s2, time.Now().Format("2006-01-02"), tt)
		routesCache[q] = routes
		saveCache()
	}

	if len(routes) > 0 {
		// Нарисовать диаграмму и сделать миниатюру с подписью
		img := drawDiagram(s1,s2,routes)
		smallImg := transform.Resize(img, 320, 180, transform.Linear)

		diagram = img
		diagramWidget.Image = smallImg
		diagramName.Text = s1 + " — " + s2
		diagramBox.Show()
	} else {
		// Если рейсов нет, то миниатюру скрыть
		diagramBox.Hide()
		dialog.ShowInformation("", "Нет рейсов", W)
		buttonBack(0)
	}

	diagramBox.Refresh()
}