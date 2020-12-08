package src

import (
	"fyne.io/fyne/dialog"
	"time"
)

// Событие нажатия кнопки
func mainButtonClick(s1,s2,tt string) {
	if s1 != s2 {
		routes := getAllSortedRoutes(s1,s2, time.Now().Format("2006-01-02"), tt)
		if len(routes) > 0 {

		} else {
			dialog.ShowInformation("", "Нет рейсов", W)
		}
	} else {
		dialog.ShowInformation("Ошибка", "Станции должны быть разными", W)
	}
}