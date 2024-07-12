package ascii

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

// Checkfiles compares local banner files to the ones in the cloud and updates them if necessary.
func Checkfiles(s string) error {
	remoteURL := ""
	switch s {
	case "resources/standard.txt":
		remoteURL = "https://learn.zone01kisumu.ke/git/root/public/raw/branch/master/subjects/ascii-art/standard.txt"
	case "resources/thinkertoy.txt":
		remoteURL = "https://learn.zone01kisumu.ke/git/root/public/raw/branch/master/subjects/ascii-art/thinkertoy.txt"
	case "resources/shadow.txt":
		remoteURL = "https://learn.zone01kisumu.ke/git/root/public/raw/branch/master/subjects/ascii-art/shadow.txt"
	default:
		remoteURL = "https://learn.zone01kisumu.ke/git/root/public/raw/branch/master/subjects/ascii-art/standard.txt"
	}

	// Fetch the content of the remote file
	resp, err := http.Get(remoteURL)
	if err != nil {
		fmt.Println("Error fetching remote file:", err)
		return err
	}
	defer resp.Body.Close()

	// Check if the request was successful
	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Error: Received non-200 response code %d for URL %s\n", resp.StatusCode, remoteURL)
		return fmt.Errorf("received non-200 response code %d for URL %s", resp.StatusCode, remoteURL)
	}

	// Read the remote file content into a string
	remoteContent, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading remote file content:", err)
		return err
	}

	// Normalize line endings from Windows-style (\r\n) to Unix-style (\n)
	normalizedContent := strings.ReplaceAll(string(remoteContent), "\r\n", "\n")

	// Remove the existing file if it exists
	if _, err := os.Stat(s); err == nil {
		err = os.Remove(s)
		if err != nil {
			fmt.Println("Error removing existing file:", err)
			return err
		}
	}

	// Create the new file
	localFile, err := os.Create(s)
	if err != nil {
		fmt.Println("Error creating new file:", err)
		return err
	}
	defer localFile.Close()

	// Write the normalized content to the new local file
	_, err = io.WriteString(localFile, normalizedContent)
	if err != nil {
		fmt.Println("Error writing content to new file:", err)
		return err
	}

	fmt.Printf("Successfully updated %s.Please try again\n", s)
	return nil
}
