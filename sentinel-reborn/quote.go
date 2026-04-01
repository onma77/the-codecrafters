package main

import (
	"regexp"
)

func fixQuotes(text string) string {
	p1 := regexp.MustCompile(`'\s+`)
	p2 := regexp.MustCompile(`\s+'`)

	text = p1.ReplaceAllString(text, "'")
	text = p2.ReplaceAllString(text, "'")
	return text
}
