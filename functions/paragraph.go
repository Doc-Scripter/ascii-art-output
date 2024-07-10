package ascii

import (
	"strings"
)

// Paragraph splits input paragraph into lines and prints the ascii art of each line
func Paragraph(input string, asciiMap map[rune][]string) []string {
	if input == "" {
		return []string{}
	}
	input = strings.Replace(input, "\n", "\\n", -1)
	var ToColor []string

	words := strings.Split(input, "\\n")
	for i, v := range words {
		if words[i] == "" {
			continue
		} else {
			ascii := Art(v, asciiMap)
			ToColor = append(ToColor, ascii)
		}
	}
	return ToColor
}
