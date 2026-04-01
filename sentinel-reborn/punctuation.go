package main

import (
	"regexp"
	"strings"
)

func fixPunctuation(text string) string {

	p1 := regexp.MustCompile(`\s+([?.,:;!]+)`)

	text = p1.ReplaceAllString(text, "$1 ")
	result := strings.Fields(text)
	return strings.Join(result, " ")
}
