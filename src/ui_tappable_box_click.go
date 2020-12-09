package src

import (
	"fyne.io/fyne"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
)

func deployResult() {
	w := A.NewWindow(" ")

	ci := &canvas.Image{
		Image:     resultImage,
		FillMode:  canvas.ImageFillOriginal,
		ScaleMode: canvas.ImageScalePixels,
	}
	b := &widget.Button{
		Text:     "",
		Style:    widget.PrimaryButton,
		Icon:     theme.ConfirmIcon(),
		OnTapped: func() {
			w.Close()
		},
	}

	vbox := widget.NewVBox(
		ci,
		widget.NewHBox(layout.NewSpacer(),b,layout.NewSpacer()),
	)

	C := fyne.NewContainerWithLayout(layout.NewCenterLayout(), vbox)
	w.SetContent(C)
	w.SetFullScreen(true)
	w.Show()
}