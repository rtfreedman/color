package color

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
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

var rgbMatch *regexp.Regexp = regexp.MustCompile(`[0-9a-fA-F]{6}`)

func generateRGBColorCode(color string) (colorCode string, err error) {
	match := rgbMatch.FindString(color)
	if match == "" {
		err = errors.New("no match found. please supply 6 digit hex code to use COLORRGB color mode")
		return
	}
	r, err := strconv.ParseInt(match[0:2], 16, 16)
	if err != nil {
		return
	}
	g, err := strconv.ParseInt(match[2:4], 16, 16)
	if err != nil {
		return
	}
	b, err := strconv.ParseInt(match[4:6], 16, 16)
	if err != nil {
		return
	}
	colorCode = fmt.Sprintf("\x1b[38;2;%d;%d;%dm", r, g, b)
	return
}

func retrieveColorCode(color string) (colorCode string, err error) {
	var ok bool
	if COLORMODE == COLOR16 {
		colorCode, ok = color16Map[strings.ToLower(color)]
	} else if COLORMODE == COLOR256 {
		colorCode, ok = color256Map[strings.ToLower(color)]
	} else if COLORMODE == COLORRGB {
		colorCode, err = generateRGBColorCode(color)
	} else { // otherwise we won't print a color and we generate an error
		err = errors.New("COLORMODE improperly set")
	}
	if !ok && err == nil {
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
