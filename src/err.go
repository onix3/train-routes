package src

import (
	"errors"
	"fmt"
	"fyne.io/fyne/dialog"
	"os"
	"strings"
	"time"
)

// Обработка
func IsErr(err error) {
	if err != nil {
		doErr(err)
	}
}

// Открытие лог-файла, запись сообщения об ошибке, закрытие лог-файла
func doErr(err error) {
	name := "ERRORS-train-timetable.txt"
	logFile, fileErr := os.OpenFile(name, os.O_APPEND, os.ModeAppend)
	if fileErr != nil {
		logFile,_ = os.Create(name)
	}

	s := fmt.Sprintf("[%s]   %s", time.Now().Format("2006.01.02 15:04:05"), err.Error())
	_,_ = fmt.Fprintln(logFile, s)
	fmt.Println(s)

	if len(err.Error()) < 140 {
		// создать ошибку, у которой сообщение wrapped
		e := errors.New(wrap(err.Error(), 50))
		// отобразить диалог с ошибкой
		dialog.ShowError(e, W)
	}

	fmt.Printf("\a") // звуковой сигнал

	_ = logFile.Close()
}

func wrap(s string, w int) string {
	var a []string
	for i:=0; i<=len(s)/w; i++ {
		start := i*w
		end := (i+1)*w
		if end > len(s) {
			end = i*w + len(s)%w
		}
		a = append(a,s[start:end])
	}
	return strings.Join(a,"\n")
}