package main

import (
	"github.com/rtfreedman/color"
)

func main() {
	color.Println("green", "16 Bit Printing!")
	color.COLORMODE = color.COLOR256
	color.Println("dark sea green", "256 Bit Printing!")
	color.COLORMODE = color.COLORRGB
	color.Println("FF3300", "It works with rgb too!")
}
