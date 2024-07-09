package main

import (
	"crypto/sha256"
	"fmt"
	"io"
	"log"
	"os"

	asciiArt "ascii/functions"
)

func main() {
	args := len(os.Args)
	if args < 2 || args > 3 {
		fmt.Println("Usage: go run . [STRING] [BANNER]\n\nEX: go run . something standard")
		return
	}

	standardCheckSum := "e194f1033442617ab8a78e1ca63a2061f5cc07a3f05ac226ed32eb9dfd22a6bf"
	shadowCheckSum := "26b94d0b134b77e9fd23e0360bfd81740f80fb7f6541d1d8c5d85e73ee550f73"
	thinkertoyCheckSum := "092d0cde973bfbb02522f18e00e8612e269f53bac358bb06f060a44abd0dbc52"
	

	// if only the string to be printed is provided
	if args == 2 {
		f, err := os.Open("resources/standard.txt")
		if err != nil {
			fmt.Print("Unable to read file.")
			return
		}
		defer f.Close()

		h := sha256.New()
		if _, err := io.Copy(h, f); err != nil {
			log.Fatal(err)
		}
		checkSum := string(fmt.Sprintf("%x", h.Sum(nil)))
		if checkSum != standardCheckSum {
			fmt.Println("The file has been modified. Please download the latest version from github.")
			return
		}

		s, err := os.ReadFile("resources/standard.txt")
		if err != nil {
			fmt.Println("File not found")
			return
		}
		m := asciiArt.AsciiArt(string(s))
		res := asciiArt.Tab(os.Args[1])
		asciiArt.Paragraph(res, m)
	}

	// if the string to be printed is provided and also the bannerfile
	if args == 3 {
		file := os.Args[2]

		banners := []string{"standard", "thinkertoy", "shadow"}
		for i := range banners {
			if file != banners[i] && i == len(banners)-1 {
				fmt.Println("Usage: go run . [STRING] [BANNER]\n\nEX: go run . something standard")
				return
			} else if file == banners[i] {
				break
			}
		}
		switch file {
		case "standard":
			file = "resources/standard.txt"
		case "thinkertoy":
			file = "resources/thinkertoy.txt"
		case "shadow":
			file = "resources/shadow.txt"
		case "ac":
			file = "resources/ac.txt"
		default:
			file = "resources/standard.txt"
		}

		f, err := os.Open(file)
		if err != nil {
			fmt.Print("Unable to read file.")
			return
		}
		defer f.Close()

		h := sha256.New()
		if _, err := io.Copy(h, f); err != nil {
			log.Fatal(err)
		}
		checkSum := string(fmt.Sprintf("%x", h.Sum(nil)))
		
		if checkSum != standardCheckSum && checkSum != thinkertoyCheckSum && checkSum != shadowCheckSum {
			fmt.Println("File contents have been corrupted. Redownloading the banner file")
			asciiArt.Checkfiles(file)
			return
		}

		s, err := os.ReadFile(file)
		if err != nil {
			fmt.Println("File not found")
			return
		}

		m := asciiArt.AsciiArt(string(s))
		res := asciiArt.Tab(os.Args[1])
		asciiArt.Paragraph(res, m)

	}
}
