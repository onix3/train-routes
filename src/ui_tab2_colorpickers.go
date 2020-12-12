package src

import (
	"fmt"
	"fyne.io/fyne"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
	"github.com/lusingander/colorpicker"
	"image/color"
)

type colorSelector struct {
	entry        *widget.Entry
	rect         colorpicker.PickerOpenWidget
	tmp          color.Color
	update       func(color.Color)
	sampleWidget fyne.CanvasObject
}

func newColorSelector(defaultColor color.Color, update func(color.Color)) *colorSelector {
	entry := &widget.Entry{}
	rect := colorpicker.NewColorSelectModalRect(W, fyne.NewSize(20, 20), defaultColor)
	selector := &colorSelector{
		entry:  entry,
		rect:   rect,
		tmp:    defaultColor,
		update: update,
	}
	selector.setColor(defaultColor)
	rect.SetOnChange(selector.setColorKeepAlpha)
	entry.OnChanged = func(s string) {
		l := len(s)
		if l == 7 {
			var r,g,b uint8
			_,err := fmt.Sscanf(s, "#%02x%02x%02x", &r, &g, &b);
			if err == nil {
				selector.setColor(color.RGBA{r, g, b, 255})
			}
		}
	}
	return selector
}

func (c *colorSelector) setColorKeepAlpha(clr color.Color) {
	r, g, b, _ := clr.RGBA()
	c.setColor(color.RGBA{R: uint8(r), G: uint8(g), B: uint8(b), A: 255})
}

func (c *colorSelector) setColor(clr color.Color) {
	c.tmp = clr
	c.entry.SetText(hexColorString(clr))
	c.rect.SetColor(clr)
	c.update(clr)
	applyTheme()
}

func hexColorString(c color.Color) string {
	rgba := color.RGBAModel.Convert(c).(color.RGBA)
	return fmt.Sprintf("#%.2X%.2X%.2X", rgba.R, rgba.G, rgba.B)
}

func colorRow(title string, defaultColor color.Color, update func(color.Color)) []fyne.CanvasObject {
	cs := newColorSelector(defaultColor, update)
	return []fyne.CanvasObject{widget.NewLabel(title), widget.NewHBox(cs.rect, cs.entry)}
}

func colorRows(t *userTheme) fyne.CanvasObject {
	all := []fyne.CanvasObject{}
	all = append(all, colorRow("Фон", t.BackgroundColor(), t.SetBackgroundColor)...)
	all = append(all, colorRow("Текст", t.TextColor(), t.SetTextColor)...)
	all = append(all, colorRow("Виджеты", t.ButtonColor(), t.SetButtonColor)...)
	all = append(all, colorRow("Акцент", t.PrimaryColor(), t.SetPrimaryColor)...)
	all = append(all, colorRow("Иконки", t.IconColor(), t.SetIconColor)...)
	all = append(all, colorRow("Бегунок", t.ScrollBarColor(), t.SetScrollBarColor)...)
	all = append(all, colorRow("Ссылка", t.HyperlinkColor(), t.SetHyperlinkColor)...)
	all = append(all, colorRow("При наведении", t.HoverColor(), t.SetHoverColor)...)

	return fyne.NewContainerWithLayout(
		layout.NewHBoxLayout(),
		fyne.NewContainerWithLayout(
			layout.NewGridLayoutWithColumns(2),
			all...,
		),
	)
}