package src

import (
	"github.com/fogleman/gg"
	"github.com/onix3/train-routes/resource"
	"image"
	"math"
	"os"
	"strconv"
	"time"
)

// Ресурс (представленный как массив байтов) преобразуется в файл,
// который появляется в папке с исполняемым файлом
func resourceToFile(data []byte, name string) {
	f,err := os.Create(name)
	IsErr(err)
	_,err = f.Write(data)
	IsErr(err)
	err = f.Close()
	IsErr(err)
}

// Диаграмма включает в себя все рейсы
// Последний рейс может закончиться на следующие сутки далеко за полночь
// То есть диапазон, охватываемый диаграммой, может составить, допустим, 24+7 = 31 час
// В этой функции вычисляется процент: "положение" времени t в этом диапазоне 0-31 часов
func timePercent(t time.Time, hoursRange int) float64 {
	d := t.Sub(today0000)
	p := d.Minutes()/float64(hoursRange*60)
	return p
}

// Получение изображения с нарисованной диаграммой
func drawDiagram(routes []route, fileName string) image.Image {
	today0000 = thisDay0000(routes[0].T1)

	W,H := 1600,900
	slr,sb := 100.0,35.0 // отступы слева-справа, снизу
	h := math.Round((float64(H)-sb)/float64(len(routes)))

	// два узора для каждого направления
	forwardPattern = pattern(h/2,h,true)
	reversePattern = pattern(h/2,h,false)

	// создание полотна и заливка
	C := gg.NewContext(W, H)
	C.SetRGBA255(128,128,128,255)
	C.Clear()

	// создать файл шрифта из ресурса и применить его
	resourceToFile(resource.ConsolasTtf.StaticContent,"consolas.ttf")
	err := C.LoadFontFace("consolas.ttf",24)
	IsErr(err)

	// сколько часов охватывает диаграмма
	hours := int(routes[len(routes)-1].T2.Sub(today0000).Hours()+1)
	if hours < 24 {
		hours = 24
	}
	// ширина, выделенная на один час на диаграмме
	oneHourWidth := (float64(W)-2*slr)/float64(hours)

	// если рейсы выполняются на следующие сутки, то обозначить сутки этот диапазон темнее
	if hours > 24 {
		C.SetRGB255(120,120,120)
		C.DrawRectangle(slr + oneHourWidth*24.0, 0,
			oneHourWidth*float64(hours-24),float64(H))
		C.Fill()
	}

	// положение каждого часа (линия и подпись)
	for i:=0; i<=hours; i++ {
		x := slr + oneHourWidth*float64(i)
		C.SetRGB255(112,112,112)
		C.DrawLine(x+0.5,0.5,x+0.5,float64(H)-0.5)
		C.Stroke()
		if i != hours {
			C.SetRGB255(128,64,64)
			C.DrawStringAnchored(strconv.Itoa(i%24),
				x+5,float64(H)-5,0,0)
		}
	}

	// ..и каждого шестого часа
	C.SetRGB255(96,96,96)
	for i:=0; i<=hours/4; i++ {
		x := math.Round(slr+ oneHourWidth*6*float64(i))
		C.DrawLine(x+0.5,0.5,x+0.5,float64(H)-0.5)
		C.Stroke()
	}

	// рейс представлен как прямоугольник и узором, характеризующим направление.
	// время отправления и прибытия отображаются слева и справа при любом направлении
	C.SetLineWidth(3)
	for i,t := range routes {
		x1 := math.Round(slr+(float64(W)-2*slr)*timePercent(t.T1,hours))+0.5
		x2 := math.Round(slr+(float64(W)-2*slr)*timePercent(t.T2,hours))+0.5
		y1 := math.Round(float64(i)*h)+0.5

		if t.S1 == select1.Selected {
			C.SetFillStyle(forwardPattern)
		} else {
			C.SetFillStyle(reversePattern)
		}
		C.DrawRectangle(x1,y1,x2-x1,h)
		C.Fill()

		C.DrawRectangle(x1,y1,x2-x1,h)
		//C.SetRGB255(80,80,80)    // width = 1
		C.SetRGB255(96,96,96) // width = 3
		C.Stroke()

		timeString1, timeString2 := t.T1.Format("15:04"),t.T2.Format("15:04")

		C.SetRGB255(255,192,0)
		C.DrawStringAnchored(timeString1,
			x1-10,y1+h/2,1.0,0.4)
		C.DrawStringAnchored(timeString2,
			x2+10,y1+h/2,0,0.4)
	}

	// применить шрифт снова (другого размера) и удалить файл
	err = C.LoadFontFace("consolas.ttf",48)
	IsErr(err)
	err = os.Remove("consolas.ttf")
	IsErr(err)

	// названия городов
	C.SetRGBA255(0,0,0,64)
	C.DrawStringAnchored(select1.Selected,
		slr+20,float64(H/2),0,0.5)
	C.DrawStringAnchored(select2.Selected,
		float64(W)-slr-20,float64(H/2),1,0.5)


	// сохранение диаграммы на Рабочий Стол
	dir,err := os.UserHomeDir()
	IsErr(err)
	err = C.SavePNG(dir + "\\Desktop\\••• " + fileName + ".png")
	IsErr(err)

	return C.Image()
}