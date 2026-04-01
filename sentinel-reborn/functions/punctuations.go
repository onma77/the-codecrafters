package main

import (
	"fmt"
	"strings"
)

func FormatPunctuation(text string) string {
	text = strings.ReplaceAll(text, " ... ", "... ")
	text = strings.ReplaceAll(text, " .", ". ")
	text = strings.ReplaceAll(text, " ,", ", ")
	text = strings.ReplaceAll(text, " !", "! ")
	text = strings.ReplaceAll(text, " !", "!")
	text = strings.ReplaceAll(text, " ?", "?")
	text = strings.ReplaceAll(text, " :", ": ")
	text = strings.ReplaceAll(text, " ;", ";")
	text = strings.ReplaceAll(text, " ' ", "'")
	return text
}
func main() {
	fmt.Println(FormatPunctuation(" Punctuation tests are ... kinda boring ,what do you think ?"))
}
