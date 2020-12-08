package src

import "github.com/fogleman/gg"

// Узор прямоугольника в диаграмме, в зависимости от направления рейса
// Прямое направление - зелёный со стрелками вправо
// Обратное направление - синий со стрелками влево
func pattern(w,h float64, isForward bool) gg.Pattern {
	C := gg.NewContext(int(w),int(h))
	C.DrawRectangle(0,0,w,h)
	if isForward {
		C.SetRGBA255(96,192,96,255)
	} else {
		C.SetRGBA255(128,128,255,255)
	}
	C.Fill()

	C.SetRGBA255(0,0,0,32)
	if isForward {
		C.DrawLine(w/4,0,w*3/4,h/2) //   \
		C.DrawLine(w*3/4,h/2,w/4,h)    //   /
	} else {
		C.DrawLine(w*3/4,0,w/4,h/2) //   /
		C.DrawLine(w/4,h/2,w*3/4,h)    //   \
	}
	C.SetLineWidth(w/2)
	C.Stroke()

	return gg.NewSurfacePattern(C.Image(), gg.RepeatBoth)
}