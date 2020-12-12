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
	"time"
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
	for i:=1; i<=17; i++ {
		numbers = append(numbers, &canvas.Text{
			Color: color.RGBA{204, 153, 102, 255},
			Text: "  " + strconv.Itoa(i) + "  ",
			Alignment: fyne.TextAlignCenter,
			TextSize: 14,
		})
	}
	scroller := widget.NewHScrollContainer(widget.NewHBox(numbers...))

	label1 := widget.NewLabel("Original theme editor is")
	label1.TextStyle.Italic = true
	parsed,_ := url.Parse("https://github.com/lusingander/fyne-theme-generator")
	hyperLink := widget.NewHyperlink("here", parsed)

	widgets := widget.NewVBox(
		widget.NewHBox(button, layout.NewSpacer(), info, confirm),
		scroller,
		widget.NewHBox(label1, hyperLink),
	)

	item := widget.NewAccordionItem("Widgets", widgets)
	accordion := widget.NewAccordionContainer(item)

	// раскрыть, чтобы окно приняло размер с учётом развёрнутого списка
	accordion.Open(0)
	// а затем закрыть
	go func() {
		time.Sleep(100*time.Millisecond)
		accordion.CloseAll()
	}()

	vbox := widget.NewVBox(
		widget.NewLabel(""),
		widget.NewHBox(layout.NewSpacer(), colorRows(T), layout.NewSpacer()),
		widget.NewLabel(""),
		widget.NewHBox(layout.NewSpacer(), accordion, layout.NewSpacer()),
		widget.NewLabel(""),
	)

	return vbox
}