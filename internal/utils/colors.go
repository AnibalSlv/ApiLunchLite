package utils

import "github.com/fatih/color"

var (
	cStop   = color.BgRGB(97, 0, 0).Add(color.FgHiRed)
	cRun    = color.BgRGB(6, 64, 43).Add(color.FgHiGreen)
	cRed    = color.New(color.FgRed)
	cBlue   = color.New(color.FgBlue)
	cGreen  = color.New(color.FgGreen)
	cYellow = color.New(color.FgYellow)
)

func Stop(txt string) string   { return cStop.Sprint(txt) }
func Run(txt string) string    { return cRun.Sprint(txt) }
func Red(txt string) string    { return cRed.Sprint(txt) }
func Blue(txt string) string   { return cBlue.Sprint(txt) }
func Green(txt string) string  { return cGreen.Sprint(txt) }
func Yellow(txt string) string { return cYellow.Sprint(txt) }
