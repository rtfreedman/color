package main

import "github.com/rtfreedman/color"

func main() {
	color.Println("green", "look it works!")
	color.COLORMODE = color.COLORRGB
	color.Println("FF3300", "in rgb too!")
}
