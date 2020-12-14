package src

import (
	"fyne.io/fyne"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/driver/desktop"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
	"net/url"
)

func deployMap() {
	w := A.Driver().(desktop.Driver).CreateSplashWindow()
	w.Resize(fyne.NewSize(ScreenWidth,ScreenHeight))
	w.CenterOnScreen()

	label1 := widget.NewLabel("Открыть")
	parsed,_ := url.Parse("https://yandex.ru/maps/?um=constructor%3Ad7846cc6cf6516763b73d7d45ca4bff3188045e89037947249cd0afb2dad6f6d&source=constructorLink")
	hyperLink := widget.NewHyperlink("карту", parsed)
	label2 := widget.NewLabel("в браузере")

	ci := &canvas.Image{
		Image:     mapImage,
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

	C := fyne.NewContainerWithLayout(layout.NewCenterLayout(),
		widget.NewVBox(
			widget.NewHBox(layout.NewSpacer(),label1,hyperLink,label2,layout.NewSpacer()),
			ci,
			widget.NewHBox(layout.NewSpacer(),b,layout.NewSpacer()),
		),
	)
	w.SetContent(C)
	w.Show()
}

func deployDiagram() {
	w := A.Driver().(desktop.Driver).CreateSplashWindow()
	w.Resize(fyne.NewSize(ScreenWidth,ScreenHeight))
	w.CenterOnScreen()

	label := widget.NewLabel("Данные предоставлены сервисом Яндекс.Расписания")
	parsed,_ := url.Parse("https://rasp.yandex.by/")
	hyperLink := widget.NewHyperlink("rasp.yandex.by", parsed)
	yandBox := widget.NewHBox(layout.NewSpacer(),label,hyperLink,layout.NewSpacer())

	ci := &canvas.Image{
		Image:     diagram,
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
	// Для первой диаграммы добавлять надпись про Янд-кс
	if СколькоДиаграммПоказано <= 0 {
		vbox.Prepend(yandBox)
	}

	C := fyne.NewContainerWithLayout(layout.NewCenterLayout(), vbox)

	w.SetContent(C)
	KeyBindings(w)
	w.Show()

	if СколькоДиаграммПоказано <= 0 {
		firstDiagramHelp(w)
	}

	СколькоДиаграммПоказано++
	saveCountOfCompletedDiagrams()
}