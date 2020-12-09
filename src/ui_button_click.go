package src

import (
	"fmt"
	"fyne.io/fyne/dialog"
	"github.com/anthonynsimon/bild/transform"
	"time"
)

// Событие нажатия кнопки
func mainButtonClick(s1,s2,tt string) {
	if s1 != s2 {
		routes := getAllSortedRoutes(s1,s2, time.Now().Format("2006-01-02"), tt)
		if len(routes) > 0 {
			// Нарисовать диаграмму и сделать миниатюру с подписью
			img := drawDiagram(routes, s1 + " → " + s2)
			smallImg := transform.Resize(img, 320, 180, transform.Linear)

			СколькоДиаграммСделано++
			saveCountOfCompletedDiagrams()
			fmt.Println("Сделано", СколькоДиаграммСделано)

			resultImage = img
			resultImageWidget.Image = smallImg
			resultText.Text = s1 + " — " + s2
			resultBox.Show()
		} else {
			// Если рейсов нет, то миниатюру скрыть
			resultBox.Hide()
			dialog.ShowInformation("", "Нет рейсов", W)
		}
	} else {
		resultBox.Hide()
		dialog.ShowInformation("Ошибка", "Станции должны быть разными", W)
		resultText.Text = ""
	}
	resultBox.Refresh()
}