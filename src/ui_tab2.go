package src

import (
	"fyne.io/fyne"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/dialog"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
	"image/color"
	"net/url"
	"strconv"
)

// Соджержимое вкладки "Тема"
func content2() fyne.CanvasObject {
	button := &widget.Button{
		Text:              "Primary",
		Style:             widget.PrimaryButton,
		Icon:              theme.ViewRefreshIcon(),
	}
	info := widget.NewButton("Info", func() {
		dialog.NewInformation("Info", "information dialog...", W)
	})
	confirm := widget.NewButton("Confirm", func() {
		dialog.NewConfirm("Confirm", "confirm dialog...", func(bool) {}, W)
	})

	var numbers []fyne.CanvasObject
	for i:=1; i<=77; i++ {
		numbers = append(numbers, &canvas.Text{
			Color: color.RGBA{204, 153, 102, 255},
			Text: "  " + strconv.Itoa(i) + "  ",
			Alignment: fyne.TextAlignCenter,
			TextSize: 14,
		})
	}
	scroller := widget.NewHScrollContainer(widget.NewHBox(numbers...))

	//selects := &widget.Select{
	//	BaseWidget:  widget.BaseWidget{},
	//	Selected:    "A",
	//	Options:     []string{"A","B","C","D","E"},
	//	PlaceHolder: "Выбирай",
	//}

	label1 := widget.NewLabel("Original theme editor is")
	label1.TextStyle.Italic = true
	parsed,_ := url.Parse("https://github.com/lusingander/fyne-theme-generator")
	hyperLink := widget.NewHyperlink("here", parsed)

	group := widget.NewGroup(
		"Виджеты для демонстрации",
		widget.NewVBox(
			//selects,
			widget.NewHBox(button, layout.NewSpacer(), info, confirm),
			scroller,
		),
	)

	vbox := widget.NewVBox(
		widget.NewLabel(""),
		widget.NewHBox(layout.NewSpacer(), colorRows(T),layout.NewSpacer()),
		widget.NewLabel(""),
		widget.NewLabel(""),
		group,
		widget.NewHBox(label1, hyperLink),
		widget.NewLabel(""),
	)

	C := fyne.NewContainerWithLayout(layout.NewCenterLayout(), vbox)
	return C
}