package main

import (
	"strconv"
	"strings"
)

func ConvToHex(num string) string {
	result, err := strconv.ParseInt(num, 16, 64)
	if err != nil {
		return num
	}
	return strconv.FormatInt(result, 10)
}
func ConvToBin(num string) string {
	result, err := strconv.ParseInt(num, 2, 64)
	if err != nil {
		return num
	}
	return strconv.FormatInt(result, 10)
}

func strConvert(text string) string {
	input := strings.Fields(text)
	for i := 0; i < len(input); i++ {
		if input[i] == "(hex)" {
			input[i-1] = ConvToHex(input[i-1])
			input = append(input[:i], input[i+1:]...)
		}

		if input[i] == "(bin)" {
			input[i-1] = ConvToBin(input[i-1])
			input = append(input[:i], input[i+1:]...)
		}

	}
	return strings.Join(input, " ")
}
