package main

import (
	"strconv"
	"strings"
)

func conv(text string) string {
	word := text
	car := ""
	man := ""
	word = strings.ReplaceAll(word, "(up, ", "(up,")
	word = strings.ReplaceAll(word, "(cap, ", "(cap,")
	word = strings.ReplaceAll(word, "(low, ", "(low,")
	words := strings.Fields(word)
	for i := 0; i < len(words); i++ {
		if strings.Contains(words[i], "(up,") {
			j := i
			car = words[j]
			for _, r := range car {
				if r >= '0' && r <= '9' {
					man = string(r)
					nums, _ := strconv.Atoi(man)
					k := j - 1
					if k+1 < nums {
						nums = k + 1
					}
					for k = j - 1; k >= j-nums; k-- {
						words[k] = strings.ToUpper(words[k])
					}
				}
			}
			words = append(words[:j], words[j+1:]...)
			i--
		}
		if strings.Contains(words[i], "(low,") {
			j := i
			car = words[j]
			for _, r := range car {
				if r >= '0' && r <= '9' {
					man = string(r)
					nums, _ := strconv.Atoi(man)
					k := j - 1
					if k+1 < nums {
						nums = k + 1
					}
					for k := j; k >= j-nums; k-- {
						words[k] = strings.ToLower(words[k])
					}
				}
			}
			words = append(words[:j], words[j+1:]...)
			i--
		}
		if strings.Contains(words[i], "(cap,") {
			j := i
			car = words[j]
			for _, r := range car {
				if r >= '0' && r <= '9' {
					man = string(r)
					nums, _ := strconv.Atoi(man)
					k := j - 1
					if k+1 < nums {
						nums = k + 1
					}
					for k := j; k >= j-nums; k-- {
						words[k] = strings.ToUpper(string(words[k][0])) + strings.ToLower(string(words[k][1:]))
					}
				}
			}
			words = append(words[:j], words[j+1:]...)
			i--
		}
	}
	return strings.Join(words, " ")

}
