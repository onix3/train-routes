package src

import (
	"encoding/json"
	"image/color"
	"strings"
)

// Применить текущую тему
func applyTheme() {
	A.Settings().SetTheme(T)
	SaveRequest = true
}

// Структура темы специально для сохранения в json
type userThemeJson struct {
	BackgroundC     color.RGBA `json:"background_c"`
	ButtonC         color.RGBA `json:"button_c"`
	TextC           color.RGBA `json:"text_color"`
	IconC           color.RGBA `json:"icon_c"`
	HyperlinkC      color.RGBA `json:"hyperlink_c"`
	PrimaryC        color.RGBA `json:"primary_c"`
	HoverC          color.RGBA `json:"hover_c"`
	ScrollBarC      color.RGBA `json:"scroll_bar_c"`
}

// Загрузить тему из preferences
func LoadTheme() {
	T = newTheme()
	var t userThemeJson
	if j := A.Preferences().StringWithFallback("theme", ""); j != "" {
		err := json.Unmarshal([]byte(j),&t)
		IsErr(err)
	} else {
		s := `{\"background_c\":{\"R\":68,\"G\":68,\"B\":68,\"A\":255},\"button_c\":{\"R\":48,\"G\":48,\"B\":48,\"A\":255},\"text_color\":{\"R\":255,\"G\":255,\"B\":255,\"A\":255},\"icon_c\":{\"R\":254,\"G\":254,\"B\":254,\"A\":255},\"hyperlink_c\":{\"R\":226,\"G\":143,\"B\":255,\"A\":255},\"primary_c\":{\"R\":25,\"G\":136,\"B\":116,\"A\":255},\"hover_c\":{\"R\":166,\"G\":106,\"B\":27,\"A\":255},\"scroll_bar_c\":{\"R\":67,\"G\":218,\"B\":28,\"A\":255}}`
		s = strings.Replace(s,"\\","",-1)
		err := json.Unmarshal([]byte(s),&t)
		IsErr(err)
	}
	T.SetBackgroundColor(t.BackgroundC)
	T.SetButtonColor(t.ButtonC)
	T.SetTextColor(t.TextC)
	T.SetIconColor(t.IconC)
	T.SetHyperlinkColor(t.HyperlinkC)
	T.SetPrimaryColor(t.PrimaryC)
	T.SetHoverColor(t.HoverC)
	T.SetScrollBarColor(t.ScrollBarC)
}

// Сохранить тему в preferences
func SaveTheme() {
	t := &userThemeJson{
		BackgroundC:     color.RGBAModel.Convert(T.BackgroundColor()).(color.RGBA),
		ButtonC:         color.RGBAModel.Convert(T.ButtonColor()).(color.RGBA),
		TextC:           color.RGBAModel.Convert(T.TextColor()).(color.RGBA),
		IconC:           color.RGBAModel.Convert(T.IconColor()).(color.RGBA),
		HyperlinkC:      color.RGBAModel.Convert(T.HyperlinkColor()).(color.RGBA),
		PrimaryC:        color.RGBAModel.Convert(T.PrimaryColor()).(color.RGBA),
		HoverC:          color.RGBAModel.Convert(T.HoverColor()).(color.RGBA),
		ScrollBarC:      color.RGBAModel.Convert(T.ScrollBarColor()).(color.RGBA),
	}
	data,err := json.Marshal(t)
	IsErr(err)
	A.Preferences().SetString("theme",string(data))
}