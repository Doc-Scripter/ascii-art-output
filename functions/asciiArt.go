package ascii

import (
	"strings"
)

// asciiArt maps characters to their respective Ascii art 8 lines long.
func AsciiArt(s string) map[rune][]string {
	wordArt := strings.Split(string(s), "\n")
	char := ' '
	m := make(map[rune][]string)

	// buffer is used to for lines containing the character temporarily
	var buffer []string

	len := len(wordArt)

	for a := 0; a < len; a++ {

		line := wordArt[a]

		// if loop store the content found in the first line which is nil
		if a == 0 {
			buffer = make([]string, 0)
		} else if a%9 == 0 {
			m[char] = buffer
			buffer = make([]string, 0)
			char++
		} else {
			buffer = append(buffer, line)
		}
	}
	return m
}
