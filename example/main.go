package main

import (
	"github.com/rtfreedman/color"
)

func main() {
	// default printing mode is in 16 bit
	color.Println("green", "16 Bit Printing!")
	// switching the colormode is possible through the COLORMODE variable
	color.COLORMODE = color.COLOR256
	color.Println("dark sea green", "256 Bit Printing!")
	// the third color mode uses a hex code to print in RGB
	color.COLORMODE = color.COLORRGB
	color.Println("FF3300", "It works with rgb too!")
}
