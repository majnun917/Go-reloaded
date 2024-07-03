package utils

import (
	//"fmt"
	"regexp"
	"strconv"
	"strings"
)

func HandlingVowels(str string) string {
	regexVowels := regexp.MustCompile(`\b(a|A)\s+([aeiouhAEIOUH])\s*`)
	str = regexVowels.ReplaceAllString(str, "${1}n $2")

	regexVowels = regexp.MustCompile(`\b(a|A)\s+([AEIOUH])\s*`)
	str = regexVowels.ReplaceAllString(str, "${1}N $2")
	return str
}
func HandlingPunctuation(str string) string {
	str = regexp.MustCompile(` ([,.!?;:])`).ReplaceAllString(str, "$1")
	str = regexp.MustCompile(`([,.!?;:])(\w+)`).ReplaceAllString(str, "$1 $2")
	return str
}

func HandlingMark(str string) string {
	content := strings.Fields(str)
	for i := 0; i < len(content)-1; i++ {
		if content[i] == "'" {
			content[i] = content[i] + content[i+1]
			content = append(content[:i+1], content[i+2:]...)
		}
	}
	for i := 1; i < len(content); i++ {
		if content[i] == "'" {
			content[i-1] = content[i-1] + content[i]
			content = append(content[:i], content[i+1:]...)
		}
	}
	return strings.Join(content, " ")
}

func HexBinToDecimal(decimal string) string {
	regexHex := regexp.MustCompile(`([0-9a-fA-F]+) \(hex\)`)

	regexBin := regexp.MustCompile(`([01]+) \(bin\)`)

	decimal = regexHex.ReplaceAllStringFunc(decimal, HexToDecimal)
	decimal = regexBin.ReplaceAllStringFunc(decimal, BinToDecimal)
	return decimal
}

func ContentEdit(content string) (string, error) {
	/*~~~~~~~~~~~~~~~Handling Hexa et Binary numbers~~~~~~~~~~~~~~~~~*/
	content = HexBinToDecimal(content)

	/*~~~~~~~~~~~~~~~Handling words conversion~~~~~~~~~~~~~~~~~*/
	regexWords := regexp.MustCompile(`[^\s()]+|\([^)]+\)`)
	words := regexWords.FindAllString(content, -1)
	regexConv := regexp.MustCompile(`(low|up|cap)(,\s*(\d+))?\)`)

	for i := 0; i < len(words); i++ {
		word := words[i]
		groups := regexConv.FindStringSubmatch(word)
		if groups != nil {
			convert := groups[1]
			num := 1
			if len(groups) > 3 && groups[3] != "" {
				num, _ = strconv.Atoi(groups[3])
			}
			switch convert {
			case "low":
				for j := 1; j <= num && i-j >= 0; j++ {
					words[i-j] = strings.TrimSpace(strings.ToLower(words[i-j]))
					words[i] = ""
				}
			case "up":
				for j := 1; j <= num && i-j >= 0; j++ {
					words[i-j] = strings.TrimSpace(strings.ToUpper(words[i-j]))
					words[i] = ""
				}
			case "cap":
				for j := 1; j <= num && i-j >= 0; j++ {
					words[i-j] = strings.TrimSpace(ToCapitalize(strings.ToLower(words[i-j])))
					words[i] = ""
				}
			}
		}
	}
	var cleanedWords []string
	for _, word := range words {
		if word != "" {
			cleanedWords = append(cleanedWords, word)
		}
	}

	/*~~~~~~~~~~~~~~~Handling punctuations~~~~~~~~~~~~~~~~~*/
	content = HandlingPunctuation(strings.Join(cleanedWords, " "))
	content = HandlingMark(content)

	/*~~~~~~~~~~~~~~~Handling vowels~~~~~~~~~~~~~~~~~*/
	content = HandlingVowels(content)

	return content, nil
}
