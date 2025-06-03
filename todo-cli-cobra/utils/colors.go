package utils

import "github.com/fatih/color"

var GreenPrint = color.New(color.FgGreen)
var RedPrint = color.New(color.FgRed)
var SuccessPrint = color.New(color.FgGreen).Add(color.Underline)
var ErrorPrint = color.New(color.FgRed).Add(color.Underline)
var TitlePrint = color.New(color.FgHiYellow)
