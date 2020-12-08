package src

import (
	"fyne.io/fyne"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
)

// Содержимое первой вкладки "Расписание"
func content1() (C fyne.CanvasObject) {
	// два селектора и кнопка
	select1 = &widget.Select{
		Selected:  allCities()[1],
		Options:   allCities(),
	}
	select2 = &widget.Select{
		Selected:  allCities()[0],
		Options:   allCities(),
	}

	mainButton := &widget.Button{
		Text:     "Намутить",
		Style:    widget.PrimaryButton,
		Icon:     theme.DocumentCreateIcon(),
		OnTapped: func() {
			s1,s2 := select1.Selected, select2.Selected
			mainButtonClick(s1,s2,"train")
		},
	}

	box := widget.NewVBox(
		widget.NewHBox(select1, select2),
		mainButton,
	)
	C = widget.NewVBox(
		widget.NewHBox(layout.NewSpacer(),box,layout.NewSpacer()),
	)
	return
}