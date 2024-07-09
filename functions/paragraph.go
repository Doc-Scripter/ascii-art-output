package ascii

import (
	"fmt"
	"strings"
)

// Paragraph splits input paragraph into lines and prints the ascii art of each line
func Paragraph(sent string, m map[rune][]string) {
	if sent == "" {
		return
	}
	sent = strings.Replace(sent, "\n", "\\n", -1)

	i := 0
	c := strings.Split(sent, "\\n")
	for _, line := range c {
		if line == "" {
			i++
			if i < len(c) {
				fmt.Println()
			}
		} else {
			Art(line, m)
		}
	}
}
