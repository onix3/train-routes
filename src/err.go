package src

import (
	"fmt"
	"os"
	"time"
)

// Посредственная обработка ошибки: открытие лог-файла, запись, закрытие лог-файла
func IsErr(err error) {
	if err != nil {
		logFile, fileErr := os.OpenFile("ERRORS-train-timetable.txt", os.O_APPEND, os.ModeAppend)
		if fileErr != nil {
			logFile,_ = os.Create("ERRORS-train-timetable.txt")
		}
		logErr(logFile, err)
		_ = logFile.Close()
	}
}

// Собственно запись
func logErr(file *os.File, err error) {
	_,_ = fmt.Fprintf(file, "[%s]   %s\n",
		time.Now().Format("2006.01.02 15:04:05"), err.Error())
	fmt.Printf("\a") // звуковой сигнал
}