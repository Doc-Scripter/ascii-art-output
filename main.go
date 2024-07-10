package main

import (
	"crypto/sha256"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	asciiArt "ascii/functions"
)

func main() {
	var output string
	flag.StringVar(&output, "output", "", "store output of ascii art to a file.")
	flag.Parse()
	args := len(os.Args)
	if args < 2 || args > 4 {
		fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER] \n\nEX: go run . --output=<fileName.txt> something standard")
		return
	}
	if output != "" && strings.Contains(os.Args[1], "/")  {
		fmt.Println("Output file cannot contain a path.")
		return
	}
	
	standardCheckSum := "e194f1033442617ab8a78e1ca63a2061f5cc07a3f05ac226ed32eb9dfd22a6bf"
	shadowCheckSum := "26b94d0b134b77e9fd23e0360bfd81740f80fb7f6541d1d8c5d85e73ee550f73"
	thinkertoyCheckSum := "092d0cde973bfbb02522f18e00e8612e269f53bac358bb06f060a44abd0dbc52"

	// if the string to be printed is provided and also the bannerfile

	if args == 2 {
		if output!=""{
			fmt.Println("Usage: go run . [STRING]\n\nEX: go run . something")
			return
		}
		data, err := os.ReadFile("resources/standard.txt")
		// confirms if the file used is available
		if err != nil {
			fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER] \n\nEX: go run . --output=<fileName.txt> something standard")
			return
		}

		mapped := asciiArt.AsciiArt(string(data))
		tab := asciiArt.Tab(os.Args[1])
		word := asciiArt.Paragraph(tab, mapped)
		fmt.Print(strings.Join(word, ""))
		return
	}


	if output!="" && !strings.HasSuffix(os.Args[1],".txt"){
		fmt.Println("Output file must be a .txt file.")
		return

	}
if output!="" && !strings.HasPrefix(os.Args[1],"--output="){
	fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER] \n\nEX: go run . --output=<fileName.txt> something standard")
	return
}

	if args == 3 {
		
		if output != ""  {

			f, err := os.Open("resources/standard.txt")
			if err != nil {
				fmt.Print("Unable to read file.")
				return
			}
			

			h := sha256.New()
			if _, err := io.Copy(h, f); err != nil {
				log.Fatal(err)
			}
			checkSum := string(fmt.Sprintf("%x", h.Sum(nil)))

			if checkSum != standardCheckSum {
				fmt.Println("File contents have been corrupted. Redownloading the banner file")
				asciiArt.Checkfiles("resources/standard.txt")
				return
			}

			s, err := os.ReadFile("resources/standard.txt")
			if err != nil {
				fmt.Println("File not found")
				return
			}

			m := asciiArt.AsciiArt(string(s))
			res := asciiArt.Tab(os.Args[2])
			result := asciiArt.Paragraph(res, m)

			err = os.WriteFile(output, []byte(strings.Join(result, "")), 0o644)
			if err != nil {
				fmt.Println("Usage: go run . [OPTION] [STRING]\n\nEX: go run . --output=<fileName.txt> something")
				return
			}

		} else{
			fmt.Println("Usage: go run . [OPTION] [STRING]\n\nEX: go run . --output=<fileName.txt> something")
			return
	}
}

	if len(os.Args) == 4 {
		
		if output != "" {
			// output := os.Args[1]
			banner := os.Args[3]

			

				banners := []string{"standard", "thinkertoy", "shadow"}
				for i := range banners {
					if banner != banners[i] && i == len(banners)-1 {
						fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER] \n\nEX: go run . --output=<fileName.txt> something standard")
						return
					} else if banner == banners[i] {
						break
					}
				}
				switch banner {
				case "standard":
					banner = "resources/standard.txt"
				case "thinkertoy":
					banner = "resources/thinkertoy.txt"
				case "shadow":
					banner = "resources/shadow.txt"
				case "ac":
					banner = "resources/ac.txt"
				default:
					banner = "resources/standard.txt"
				}

				f, err := os.Open(banner)
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
					asciiArt.Checkfiles(banner)
					return
				}

				s, err := os.ReadFile(banner)
				if err != nil {
					fmt.Println("File not found")
					return
				}

				m := asciiArt.AsciiArt(string(s))
				res := asciiArt.Tab(os.Args[1])
				result := asciiArt.Paragraph(res, m)
				err = os.WriteFile(output, []byte(strings.Join(result, "")), 0o644)
				if err != nil {
					panic(err)
				}

		}
	}
}



// func input(s []string)string {

// }