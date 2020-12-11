package src

import (
	"fmt"
	"fyne.io/fyne/dialog"
	"github.com/anthonynsimon/bild/transform"
	"time"
)

// Событие нажатия кнопки
func mainButtonClick(s1,s2,tt string) {
	q := fmt.Sprintf("%s-%s(%s)",
		s1,s2, time.Now().Format("2006-01-02"))

	var routes []route

	// кэш расписания доступен оффлайн
	// подключение в Интернету проверять, если нет в кэше
	if inCache(q) {
		routes = routesCache[q]
	} else {
		PastCity1,PastCity2 = s1,s2
		saveLastCities(s1,s2)

		if s1 == s2 {
			resultBox.Hide()
			dialog.ShowInformation("", "Города должны быть разными", W)
			resultText.Text = ""
			return
		}

		routes = getAllSortedRoutes(s1,s2, time.Now().Format("2006-01-02"), tt)
		routesCache[q] = routes
		saveCache()
	}

	if len(routes) > 0 {
		// Нарисовать диаграмму и сделать миниатюру с подписью
		img := drawDiagram(routes, s1 + " → " + s2)
		smallImg := transform.Resize(img, 320, 180, transform.Linear)

		СколькоДиаграммСделано++
		saveCountOfCompletedDiagrams()

		resultImage = img
		resultImageWidget.Image = smallImg
		resultText.Text = s1 + " — " + s2
		resultBox.Show()
	} else {
		// Если рейсов нет, то миниатюру скрыть
		resultBox.Hide()
		dialog.ShowInformation("", "Нет рейсов", W)
	}

	resultBox.Refresh()
}