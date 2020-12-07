package main

import (
	"github.com/onix3/train-routes/src"
	"time"
)

func main() {
	_ = src.GetAllSortedRoutes("Брест","Минск",time.Now().Format("2006-01-02"),"train")
}