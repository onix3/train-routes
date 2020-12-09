package src

import (
	"fmt"
	"fyne.io/fyne"
	"fyne.io/fyne/widget"
)

type tappableBox struct {
	widget.Box
}

func (t *tappableBox) Tapped(_ *fyne.PointEvent) {
	fmt.Println("Tapped")
}
func (t *tappableBox) TappedSecondary(_ *fyne.PointEvent) {
	fmt.Println("Tapped")
}

// Расширение виджета Box, чтобы обрабатывалось событие нажатия кнопки мыши
func newTappableBox(b *widget.Box) *tappableBox {
	tb := &tappableBox{
		Box: *b,
	}
	tb.ExtendBaseWidget(tb)

	return tb
}