package src

import (
	"fyne.io/fyne"
	"fyne.io/fyne/theme"
	"image/color"
)

type userTheme struct {
	backgroundColor     color.Color
	buttonColor         color.Color
	disabledButtonColor color.Color
	textColor         color.Color
	disabledTextColor color.Color
	iconColor         color.Color
	disabledIconColor color.Color
	hyperlinkColor    color.Color
	placeHolderColor  color.Color
	primaryColor      color.Color
	hoverColor        color.Color
	focusColor        color.Color
	scrollBarColor    color.Color
	shadowColor       color.Color

	textSize           int
	textFont           fyne.Resource
	textBoldFont       fyne.Resource
	textItalicFont     fyne.Resource
	textBoldItalicFont fyne.Resource
	textMonospaceFont  fyne.Resource

	padding            int
	iconInlineSize     int
	scrollBarSize      int
	scrollBarSmallSize int
}

func NewTheme() *userTheme {
	return &userTheme{
		backgroundColor:     theme.DarkTheme().BackgroundColor(),
		buttonColor:         theme.DarkTheme().ButtonColor(),
		disabledButtonColor: theme.DarkTheme().DisabledButtonColor(),
		textColor:           theme.DarkTheme().TextColor(),
		disabledTextColor:   theme.DarkTheme().DisabledTextColor(),
		iconColor:           theme.DarkTheme().IconColor(),
		disabledIconColor:   theme.DarkTheme().DisabledIconColor(),
		hyperlinkColor:      theme.DarkTheme().HyperlinkColor(),
		placeHolderColor:    theme.DarkTheme().PlaceHolderColor(),
		primaryColor:        theme.DarkTheme().PrimaryColor(),
		hoverColor:          theme.DarkTheme().HoverColor(),
		focusColor:          theme.DarkTheme().FocusColor(),
		scrollBarColor:      theme.DarkTheme().ScrollBarColor(),
		shadowColor:         theme.DarkTheme().ShadowColor(),
		textSize:            theme.DarkTheme().TextSize(),
		textFont:            theme.DarkTheme().TextFont(),
		textBoldFont:        theme.DarkTheme().TextBoldFont(),
		textItalicFont:      theme.DarkTheme().TextItalicFont(),
		textBoldItalicFont:  theme.DarkTheme().TextBoldItalicFont(),
		textMonospaceFont:   theme.DarkTheme().TextMonospaceFont(),
		padding:             theme.DarkTheme().Padding(),
		iconInlineSize:      theme.DarkTheme().IconInlineSize(),
		scrollBarSize:       theme.DarkTheme().ScrollBarSize(),
		scrollBarSmallSize:  theme.DarkTheme().ScrollBarSmallSize(),
	}
}

func (s *userTheme) BackgroundColor() color.Color      { return s.backgroundColor }
func (s *userTheme) ButtonColor() color.Color          { return s.buttonColor }
func (s *userTheme) DisabledButtonColor() color.Color  { return s.disabledButtonColor }
func (s *userTheme) TextColor() color.Color            { return s.textColor }
func (s *userTheme) DisabledTextColor() color.Color    { return s.disabledTextColor }
func (s *userTheme) IconColor() color.Color            { return s.iconColor }
func (s *userTheme) DisabledIconColor() color.Color    { return s.disabledIconColor }
func (s *userTheme) HyperlinkColor() color.Color       { return s.hyperlinkColor }
func (s *userTheme) PlaceHolderColor() color.Color     { return s.placeHolderColor }
func (s *userTheme) PrimaryColor() color.Color         { return s.primaryColor }
func (s *userTheme) HoverColor() color.Color           { return s.hoverColor }
func (s *userTheme) FocusColor() color.Color           { return s.focusColor }
func (s *userTheme) ScrollBarColor() color.Color       { return s.scrollBarColor }
func (s *userTheme) ShadowColor() color.Color          { return s.shadowColor }
func (s *userTheme) TextSize() int                     { return s.textSize }
func (s *userTheme) TextFont() fyne.Resource           { return s.textFont }
func (s *userTheme) TextBoldFont() fyne.Resource       { return s.textBoldFont }
func (s *userTheme) TextItalicFont() fyne.Resource     { return s.textItalicFont }
func (s *userTheme) TextBoldItalicFont() fyne.Resource { return s.textBoldItalicFont }
func (s *userTheme) TextMonospaceFont() fyne.Resource  { return s.textMonospaceFont }
func (s *userTheme) Padding() int                      { return s.padding }
func (s *userTheme) IconInlineSize() int               { return s.iconInlineSize }
func (s *userTheme) ScrollBarSize() int                { return s.scrollBarSize }
func (s *userTheme) ScrollBarSmallSize() int           { return s.scrollBarSmallSize }

func (s *userTheme) SetBackgroundColor(c color.Color)     { s.backgroundColor = c }
func (s *userTheme) SetButtonColor(c color.Color)         { s.buttonColor = c }
func (s *userTheme) SetDisabledButtonColor(c color.Color) { s.disabledButtonColor = c }
func (s *userTheme) SetTextColor(c color.Color)           { s.textColor = c }
func (s *userTheme) SetDisabledTextColor(c color.Color)   { s.disabledTextColor = c }
func (s *userTheme) SetIconColor(c color.Color)           { s.iconColor = c }
func (s *userTheme) SetDisabledIconColor(c color.Color)   { s.disabledIconColor = c }
func (s *userTheme) SetHyperlinkColor(c color.Color)      { s.hyperlinkColor = c }
func (s *userTheme) SetPlaceHolderColor(c color.Color)    { s.placeHolderColor = c }
func (s *userTheme) SetPrimaryColor(c color.Color)        { s.primaryColor = c }
func (s *userTheme) SetHoverColor(c color.Color)          { s.hoverColor = c }
func (s *userTheme) SetFocusColor(c color.Color)          { s.focusColor = c }
func (s *userTheme) SetScrollBarColor(c color.Color)      { s.scrollBarColor = c }
func (s *userTheme) SetShadowColor(c color.Color)         { s.shadowColor = c }