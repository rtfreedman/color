package color

import (
	"errors"
	"fmt"
)

const noColor string = "\033[0m"

// COLORMODE is initially set to COLOR16 as the default color mode (supporting the standard 16 colors)
// COLOR256 and COLORRGB are also available
var COLORMODE int = COLOR16

// Color modes
const (
	COLOR16 = iota
	COLOR256
	COLORRGB
)

func retrieveColorCode(color string) (colorCode string, err error) {
	var ok bool
	if COLORMODE == COLOR16 {
		colorCode, ok = color16Map[color]
	} else if COLORMODE == COLOR256 {
		colorCode, ok = color256Map[color]
	} else { // otherwise we won't print a color and we generate an error
		err = errors.New("COLORMODE improperly set")
	}
	if !ok && err != nil {
		err = errors.New("bad color lookup in color map")
	}
	return
}

// Print wraps the fmt.Print function handling color printing
// on error (bad color supplied, etc) the function will still print without color and return an error
func Print(color string, args ...interface{}) (n int, err error) {
	colorCode, err := retrieveColorCode(color)
	args = append(append([]interface{}{colorCode}, args...), noColor)
	n, fmtErr := fmt.Print(args...)
	// prioritize the fmt errors
	if fmtErr != nil {
		err = fmtErr
	}
	return
}

// Println wraps the fmt.Println function handling color printing
// on error (bad color supplied, etc) the function will still print without color and return an error
func Println(color string, args ...interface{}) (n int, err error) {
	colorCode, err := retrieveColorCode(color)
	args = append(append([]interface{}{colorCode}, args...), noColor)
	n, fmtErr := fmt.Println(args...)
	// prioritize the fmt errors
	if fmtErr != nil {
		err = fmtErr
	}
	return
}

// Printf wraps the fmt.Printf function handling color printing
// on error (bad color supplied, etc) the function will still print without color and return an error
func Printf(color, format string, args ...interface{}) (n int, err error) {
	colorCode, err := retrieveColorCode(color)
	args = append(append([]interface{}{colorCode}, args...), noColor)
	n, fmtErr := fmt.Printf(format, args...)
	// prioritize the fmt errors
	if fmtErr != nil {
		err = fmtErr
	}
	return
}
