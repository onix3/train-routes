package src

import (
	"syscall"
)

func getSize(nIndex int) int {
	user32 := syscall.NewLazyDLL("User32.dll")
	getSystemMetrics := user32.NewProc("GetSystemMetrics")
	index := uintptr(nIndex)
	ret, _, _ := getSystemMetrics.Call(index)
	return int(ret)
}

// Получение разрешения экрана (вертикальной составляющей)
func GetScreenHeight() int {
	unscaleHeight := getSize(1)       // 864
	scale := W.Canvas().Scale()              // 1.25
	return int(float32(unscaleHeight)*scale) // 864 * 1.25 = 1080
}


