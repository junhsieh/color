package color

import (
	"log"
)

const (
	Begin = "\033["
	End   = "\033[0m"
)

// Foreground bright
const (
	FGBlack   = "30;1"
	FGRed     = "31;1"
	FGGreen   = "32;1"
	FGYellow  = "33;1"
	FGBlue    = "34;1"
	FGMagenta = "35;1"
	FGCyan    = "36;1"
	FGWhite   = "37;1"
)

// Foreground dark
const (
	FGDarkBlack   = "30"
	FGDarkRed     = "31"
	FGDarkGreen   = "32"
	FGDarkYellow  = "33"
	FGDarkBlue    = "34"
	FGDarkMagenta = "35"
	FGDarkCyan    = "36"
	FGDarkWhite   = "37"
)

// Background bright
const (
	BGBlack   = "40;1"
	BGRed     = "41;1"
	BGGreen   = "42;1"
	BGYellow  = "43;1"
	BGBlue    = "44;1"
	BGMagenta = "45;1"
	BGCyan    = "46;1"
	BGWhite   = "47;1"
)

func printColor(format string, color string, v ...interface{}) {
	log.Printf(Begin+color+"m"+format+End, v...)
}

func Red(format string, v ...interface{}) {
	printColor(format, FGRed, v...)
}

func Green(format string, v ...interface{}) {
	printColor(format, FGGreen, v...)
}

func Yellow(format string, v ...interface{}) {
	printColor(format, FGYellow, v...)
}

func Magenta(format string, v ...interface{}) {
	printColor(format, FGMagenta, v...)
}

func Cyan(format string, v ...interface{}) {
	printColor(format, FGCyan, v...)
}
