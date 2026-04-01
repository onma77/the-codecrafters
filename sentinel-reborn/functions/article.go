package main

import (
	"strings"
)

func fixarticles(s string) string {
	words := strings.Fields(s)
	vowels := "aeiouhAEIOUH"

	for i := 0; i < len(words); i++ {
		if i < len(words) {
			word := words[i]
			if word == "a" && strings.ContainsAny(vowels, string(string(words[i+1][0]))) {
				words[i] = "an"
			} else if word == "A" && strings.ContainsAny(vowels, string(string(words[i+1][0]))) {
				words[i] = "An"
			} else if word == "an" && !strings.ContainsAny(vowels, string(string(words[i+1][0]))) {
				words[i] = "a"
			}
		}

	}

	return strings.Join(words, " ")

}
