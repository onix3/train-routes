package src

import (
	"github.com/fogleman/gg"
	"github.com/onix3/train-timetable/resource"
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

// самое позднее время прибытия
func latestArrival(routes []route) (t time.Time) {
	for _,r := range routes {
		if r.T2.After(t) {
			t = r.T2
		}
	}
	return
}

// Диаграмма включает в себя все рейсы
// Последний рейс может закончиться на следующие сутки далеко за полночь
// То есть диапазон, охватываемый диаграммой, может составить, допустим, 24+7 = 31 час
// В этой функции вычисляется процент: "положение" времени t в этом диапазоне 0-31 часов
func timePercent(t time.Time, hoursRange int) float64 {
	d := t.Sub(Now0000())
	p := d.Minutes()/float64(hoursRange*60)
	return p
}

// Получение изображения с нарисованной диаграммой
func drawDiagram(c1,c2 string, routes []route) image.Image {
	coef := 0.9
	// Для первой диаграммы оставить место для надписи про Янд-кс
	if СколькоДиаграммПоказано <= 0 {
		coef = 0.875
	}

	// Посчитать сколько займёт диаграмма при разворачивании изображения на весь экран
	H := int(float64(ScreenHeight)*coef)
	W := H*16/9
	slr,sb := 100.0,35.0 // отступы слева-справа, снизу
	h := math.Round((float64(H)-sb)/float64(len(routes)))

	// два узора для каждого направления
	forwardPattern := pattern(h/2,h,true)
	reversePattern := pattern(h/2,h,false)

	// создание полотна и заливка
	C := gg.NewContext(W, H)
	C.SetRGBA255(128,128,128,255)
	C.Clear()

	// создать файл шрифта из ресурса и применить его
	resourceToFile(resource.ConsolasTtf.StaticContent,"consolas.ttf")
	err := C.LoadFontFace("consolas.ttf",24)
	IsErr(err)

	// сколько часов охватывает диаграмма
	now0000 := Now0000()
	hours := int(latestArrival(routes).Sub(now0000).Hours()+1)
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

	// найти индекс обратного рейса, перед которым прямой рейс
	indexOfFirstReverse := 0
	for i:=1; i<len(routes); i++ {
		if c1 == routes[i-1].S1 && routes[i-1].S1 != routes[i].S1 {
			indexOfFirstReverse = i
			break
		}
	}
	// координаты области
	var rx0,ry0,rx1,ry1 int
	// каждый раз задавать пустую область
	cropRect = image.Rect(0,0,0,0)

	// рейс представлен как прямоугольник и узором, характеризующим направление.
	// время отправления и прибытия отображаются слева и справа при любом направлении
	C.SetLineWidth(3)
	for i,t := range routes {
		tp1,tp2 := timePercent(t.T1,hours),timePercent(t.T2,hours)
		x1 := math.Round(slr+(float64(W)-2*slr)*tp1)+0.5
		x2 := math.Round(slr+(float64(W)-2*slr)*tp2)+0.5
		y1 := math.Round(float64(i)*h)+0.5

		if t.S1 == c1 {
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
		C.DrawStringAnchored(timeString1, x1-10, y1+h/2,1.0,0.4)
		C.DrawStringAnchored(timeString2, x2+10, y1+h/2,0,0.4)

		if indexOfFirstReverse != 0 && i == indexOfFirstReverse-1 {
			strW,_ := C.MeasureString(timeString1)
			rx0 = int(x1-strW)-10
			ry0 = int(y1)
		}
		if indexOfFirstReverse != 0 && i == indexOfFirstReverse {
			strW,_ := C.MeasureString(timeString2)
			rx1 = int(x2+strW)+10
			ry1 = int(y1+h)
		}
	}

	if indexOfFirstReverse != 0 {
		cropRect = image.Rect(rx0-20,ry0,rx1+20,ry1)
	}

	//// Для первой диаграммы пояснить узоры
	//if СколькоДиаграммПоказано <= 0 {
	//	w,h := C.MeasureString("Обратное направление")
	//	C.DrawStringAnchored("Прямое направление", 10+w+10, float64(H)-5-h-5-h,1.0,0.4)
	//	C.DrawStringAnchored("Обратное направление", 10+w+10, float64(H)-5-h-5-h,1.0,0.4)
	//}

	// применить шрифт снова (другого размера) и удалить файл
	err = C.LoadFontFace("consolas.ttf",48)
	IsErr(err)
	err = os.Remove("consolas.ttf")
	IsErr(err)

	// названия городов
	C.SetRGBA255(0,0,0,64)
	C.DrawStringAnchored(c1, slr+20,float64(H/2),0,0.5)
	C.DrawStringAnchored(c2, float64(W)-slr-20,float64(H/2),1,0.5)

	return C.Image()
}