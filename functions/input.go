package ascii

import (
	"crypto/sha256"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

// handles command line inputs checking if correctly formatted
// checks if the banner file is modified and is not as it should be
// it downloads the banner file if different from original
// checks flags usage from CLI if has correct format
func Input(input []string) {
	var output string
	flag.StringVar(&output, "output", "", "store output of ascii art to a file.")
	flag.Parse()
	args := len(input)

	if args < 2 || args > 4 {
		fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER] \n\nEX: go run . --output=<fileName.txt> something standard")
		return
	}
	// checks if flag value is an empty string or directory path
	if output != "" && strings.Contains(os.Args[1], "/") {
		fmt.Println("Output file cannot contain a path.")
		return

	}
	// preset the checksum values of the files
	standardCheckSum := "e194f1033442617ab8a78e1ca63a2061f5cc07a3f05ac226ed32eb9dfd22a6bf"
	shadowCheckSum := "26b94d0b134b77e9fd23e0360bfd81740f80fb7f6541d1d8c5d85e73ee550f73"
	thinkertoyCheckSum := "092d0cde973bfbb02522f18e00e8612e269f53bac358bb06f060a44abd0dbc52"

	if args == 2 {
		if output != "" {
			fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER] \n\nEX: go run . --output=<fileName.txt> something standard")
			return
		}
		data, err := os.ReadFile("resources/standard.txt")
		if err != nil {
			fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER] \n\nEX: go run . --output=<fileName.txt> something standard")
			return
		}

		// the bannerfile is mapped
		mapped := AsciiArtMap(string(data))
		// handles \t controller aspect in the string
		result := Tab(os.Args[1])
		// the ascii art is then returned as slices and joined into one output
		asciiArt := AsciiCombine(result, mapped)
		fmt.Print(strings.Join(asciiArt, ""))
		return
	}

	if output != "" && !strings.HasPrefix(os.Args[1], "--output=") {
		fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER] \n\nEX: go run . --output=<fileName.txt> something standard")
		return
	}

	if args == 3 {
		if output != "" {
			if output == ".txt" {
				fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER] \n\nEX: go run . --output=<fileName.txt> something standard")
				return
			}

			f, err := os.Open("resources/standard.txt")
			if err != nil {
				fmt.Print("Unable to read file.")
				return
			}
			// a sha-256 hash object is created to store the contents of the banner file
			bannerTemp := sha256.New()
			if _, err := io.Copy(bannerTemp, f); err != nil {
				log.Fatal(err)
			}
			// the final checksum is computed and converted to hex string
			checkSum := fmt.Sprintf("%x", bannerTemp.Sum(nil))

			if checkSum != standardCheckSum {
				fmt.Println("File contents have been corrupted. Redownloading the banner file")
				Checkfiles("resources/standard.txt")
				return
			}

			bannerFile, err := os.ReadFile("resources/standard.txt")
			if err != nil {
				fmt.Println("File not found")
				return
			}

			mapped := AsciiArtMap(string(bannerFile))
			result := Tab(os.Args[2])
			asciiArt := AsciiCombine(result, mapped)

			err = os.WriteFile(output, []byte(strings.Join(asciiArt, "")), 0o644)
			if err != nil {
				fmt.Println("Usage: go run . [OPTION] [STRING]\n\nEX: go run . --output=<fileName.txt> something")
				return
			}

		} else if strings.Contains(os.Args[1], "--output=") && output == "" {
			fmt.Println("Usage: go run . [OPTION] [STRING]\n\nEX: go run . --output=<fileName.txt> something")
			return
		} else if output == "" && !strings.Contains(os.Args[1], "--output=") {
			banner := strings.ToLower(os.Args[2])
			banners := []string{"standard", "thinkertoy", "shadow"}
			for i := range banners {
				if banner != banners[i] && i == len(banners)-1 {
					fmt.Println("Usage: go run . [STRING] [BANNER] \n\nEX: go run .  something standard")
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

			bannerTemp := sha256.New()
			if _, err := io.Copy(bannerTemp, f); err != nil {
				log.Fatal(err)
			}
			checkSum := string(fmt.Sprintf("%x", bannerTemp.Sum(nil)))

			if checkSum != standardCheckSum && checkSum != thinkertoyCheckSum && checkSum != shadowCheckSum {
				fmt.Println("File contents have been corrupted. Redownloading the banner file")
				Checkfiles(banner)
				return
			}

			bannerFile, err := os.ReadFile(banner)
			if err != nil {
				fmt.Println("File not found")
				return
			}

			mapped := AsciiArtMap(string(bannerFile))
			result := Tab(os.Args[1])
			asciiArt := AsciiCombine(result, mapped)
			fmt.Print(strings.Join(asciiArt, ""))
			return
		}
	}

	if len(os.Args) == 4 {
		if output != "" {
			if output == ".txt" {
				fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER] \n\nEX: go run . --output=<fileName.txt> something standard")
				return
			}

			banner := strings.ToLower(os.Args[3])

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

			bannerTemp := sha256.New()
			if _, err := io.Copy(bannerTemp, f); err != nil {
				log.Fatal(err)
			}
			checkSum := string(fmt.Sprintf("%x", bannerTemp.Sum(nil)))

			if checkSum != standardCheckSum && checkSum != thinkertoyCheckSum && checkSum != shadowCheckSum {
				fmt.Println("File contents have been corrupted. Redownloading the banner file")
				Checkfiles(banner)
				return
			}

			bannerFile, err := os.ReadFile(banner)
			if err != nil {
				fmt.Println("File not found")
				return
			}

			mapped := AsciiArtMap(string(bannerFile))
			result := Tab(os.Args[2])
			asciiArt := AsciiCombine(result, mapped)
			err = os.WriteFile(output, []byte(strings.Join(asciiArt, "")), 0o644)
			if err != nil {
				panic(err)
			}

		} else {
			if output == "" || output == ".txt" {
				fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER] \n\nEX: go run . --output=<fileName.txt> something standard")
				return
			}
		}
	}
}
