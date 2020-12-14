package src

import (
	"fyne.io/fyne"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/dialog"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
	"github.com/anthonynsimon/bild/transform"
	"image"
)

// При первом отображении диаграммы появляется пояснительный диалог
func firstDiagramHelp(parent fyne.Window) {
	content := &widget.Box{}

	labelEnter := widget.NewLabel("Enter")
	labelEnter.TextStyle.Bold = true

	if cropRect != image.Rect(0,0,0,0) {
		img := transform.Crop(diagram, cropRect)
		//w,h := img.Bounds().Max.X, img.Bounds().Max.Y
		ci := &canvas.Image{
			Image:        img,
			FillMode:     canvas.ImageFillOriginal,
			ScaleMode:    canvas.ImageScalePixels,
		}

		content = widget.NewVBox(
			widget.NewHBox(boldL(PastCity1 + "   >>>   " + PastCity2 + "     "), layout.NewSpacer(), italicL("прямой")),
			ci,
			widget.NewHBox(boldL(PastCity1 + "   <<<   " + PastCity2 + "     "), layout.NewSpacer(), italicL("обратный")),
			widget.NewLabel(""),
			widget.NewLabel(""),
			widget.NewHBox(labelEnter, layout.NewSpacer(), widget.NewLabel("Сохранить диаграмму\nна Рабочий стол")),
			widget.NewLabel(""),
		)
	} else {
		content = widget.NewVBox(
			widget.NewHBox(labelEnter, layout.NewSpacer(), widget.NewLabel("Сохранить диаграмму\nна Рабочий стол")),
			widget.NewLabel(""),
		)
	}

	dialog.ShowCustom("", "Понятно", content, parent)
}

func boldL(s string) *widget.Label {
	l := widget.NewLabel(s)
	l.TextStyle.Bold = true
	return l
}

func italicL(s string) *widget.Label {
	l := widget.NewLabel(s)
	l.TextStyle.Italic = true
	return l
}
