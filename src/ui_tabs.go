package src

import (
	"fyne.io/fyne"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
	"strings"
)

func Content() fyne.CanvasObject {
	tabs := widget.NewTabContainer(
		&widget.TabItem{
			Text:    "Расписание",
			Icon:    theme.ContentPasteIcon(),
			Content: content1(),
		},
		&widget.TabItem{
			Text:    "Тема",
			Icon:    theme.SettingsIcon(),
			Content: widget.NewLabel(strings.Repeat(" ",77)),
		},
	)

	tabs.SetTabLocation(widget.TabLocationLeading)

	return tabs
}