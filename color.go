package color

import "fmt"

const noColor string = "\033[0m"

var colors = map[string]string{
	"black":        "\033[0;30m",
	"dark gray":    "\033[1;30m",
	"red":          "\033[0;31m",
	"light red":    "\033[1;31m",
	"green":        "\033[0;32m",
	"light green":  "\033[1;32m",
	"brown":        "\033[1;33m",
	"orange":       "\033[0;33m",
	"yellow":       "\033[1;33m",
	"blue":         "\033[0;34m",
	"light blue":   "\033[1;34m",
	"purple":       "\033[0;35m",
	"light purple": "\033[1;35m",
	"cyan":         "\033[0;36m",
	"light cyan":   "\033[1;36m",
	"light gray":   "\033[0;37m",
	"white":        "\033[1;37m",
}

// Print wraps the fmt.Print function handling color printing
// on error (bad color supplied, etc) the function will still print without color and return an error
func Print(color string, args ...interface{}) (err error) {
	colorCode := colors[color]
	args = append(append([]interface{}{colorCode}, args...), noColor)
	fmt.Print(args...)
	return
}

// Println wraps the fmt.Println function handling color printing
// on error (bad color supplied, etc) the function will still print without color and return an error
func Println(color string, args ...interface{}) (err error) {
	colorCode := colors[color]
	args = append(append([]interface{}{colorCode}, args...), noColor)
	fmt.Println(args...)
	return
}

// Printf wraps the fmt.Printf function handling color printing
// on error (bad color supplied, etc) the function will still print without color and return an error
func Printf(color, format string, args ...interface{}) (n int, err error) {
	colorCode := colors[color]
	args = append(append([]interface{}{colorCode}, args...), noColor)
	n, err = fmt.Printf(format, args...)
	return
}
