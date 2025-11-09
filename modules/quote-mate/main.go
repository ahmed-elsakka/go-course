package main

import (
	"github.com/ahmed-elsakka/go-course/modules/quote-mate/quote"
	"github.com/fatih/color"
)

func main() {
	color.Cyan("Welcome to QuoteMate!")
	color.Green(quote.GetQuote())
}
