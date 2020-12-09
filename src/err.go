package src

import (
	"fmt"
	"os"
	"time"
)

// Обработка
func IsErr(err error) {
	if err != nil {
		logErr(err)
	}
}

// Открытие лог-файла, запись, закрытие лог-файла
func logErr(err error) {
	logFile, fileErr := os.OpenFile("ERRORS-train-timetable.txt", os.O_APPEND, os.ModeAppend)
	if fileErr != nil {
		logFile,_ = os.Create("ERRORS-train-timetable.txt")
	}
	_ = logFile.Close()

	_,_ = fmt.Fprintf(logFile, "[%s]   %s\n",
		time.Now().Format("2006.01.02 15:04:05"), err.Error())
	fmt.Printf("\a") // звуковой сигнал
}