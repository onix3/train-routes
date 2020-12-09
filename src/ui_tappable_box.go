package src

import (
	"fyne.io/fyne"
	"fyne.io/fyne/widget"
)

type tappableBox struct {
	widget.Box
	deploy func()
}

func (t *tappableBox) Tapped(_ *fyne.PointEvent) {
	t.deploy()
}
func (t *tappableBox) TappedSecondary(_ *fyne.PointEvent) {
	t.deploy()
}

// Расширение виджета Box, чтобы обрабатывалось событие нажатия кнопки мыши
func newTappableBox(b *widget.Box, click func()) *tappableBox {
	tb := &tappableBox{
		Box: *b,
		deploy: click,
	}
	tb.ExtendBaseWidget(tb)

	return tb
}