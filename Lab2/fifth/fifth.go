package main

import (
	"fmt"
)

func main() {
	rect := Rectangle{Width: 5.12, Height: 2.03}
	fmt.Println("Площадь прямоугольника:", rect.Area())
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}
