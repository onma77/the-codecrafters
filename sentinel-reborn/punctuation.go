package main

import (
	"regexp"
)

func fixPunctuation(text string) string {

	p1 := regexp.MustCompile(`\s+([?.,:;!]+)`)
	p2 := regexp.MustCompile(`([?.,:;!]+)`)

	text = p1.ReplaceAllString(text, "$1")
	text = p2.ReplaceAllString(text, "$1 ")
	return text
}
