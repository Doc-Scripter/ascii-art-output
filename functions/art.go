package ascii

import (
	"fmt"
)

// Art prints a line of Ascii art as per the input word and mapping of m
func Art(text string, m map[rune][]string) {
	s := ""
	// ranges through the 8 lines which contain aspects of the ascii art
	for b := 0; b < 8; b++ {
		// ranges through the text that should be rep in ascii art
		for _, value := range text {
			graphics, ok := m[value]
			if !ok {
				fmt.Println("Ascii art not contained in file")
				return
			}
			s += graphics[b]
		}
		s += "\n"
	}
	// prints the ascii art of the string
	fmt.Print(s)
}
