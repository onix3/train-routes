package main

import (
	"errors"
	"github.com/onix3/train-routes/src"
)

func main() {
	err := errors.New("Hello, Error!")
	src.IsErr(err)
}