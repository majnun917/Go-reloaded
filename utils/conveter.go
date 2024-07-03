package utils

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func HexToDecimal(match string) string {
	hex := strings.TrimSuffix(match, " (hex)")
	decimal, err := strconv.ParseInt(hex, 16, 64)
	if err != nil {
		fmt.Println("Error converting hex to decimal:", err)
		return match
	}
	return strconv.FormatInt(decimal, 10)
}

func BinToDecimal(match string) string {
	bin := strings.TrimSuffix(match, " (bin)")
	decimal, err := strconv.ParseInt(bin, 2, 64)
	if err != nil {
		fmt.Println("Error converting bin to decimal:", err)
		return match
	}
	return strconv.FormatInt(decimal, 10)
}

func ToCapitalize(match string) string {
	match = strings.ToLower(match)
	cap := strings.TrimSuffix(match, "(cap)")
	var output []rune
	isWord := true
	for _, v := range cap {
		if isWord {
			output = append(output, unicode.ToUpper(v))
			isWord = false
		} else {
			output = append(output, v)
		}
	}
	return (string(output))
}
