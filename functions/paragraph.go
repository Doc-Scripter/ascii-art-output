package ascii

import (
	"strings"
)

// Handles newlines aspect in input string,combines all input slices' asciiArt
func AsciiCombine(input string, asciiMap map[rune][]string) []string {
	if input == "" {
		return []string{}
	}
	input = strings.Replace(input, "\n", "\\n", -1)
	var result []string

	words := strings.Split(input, "\\n")
	for i, char := range words {
		if words[i] == "" {
			continue
		} else {
			ascii := Art(char, asciiMap)
			result = append(result, ascii)
		}
	}
	return result
}
