package main

import (
	"os"

	asciiArt "ascii/functions"
)

// handles receiving and redirecting  command line inputs
func main() {
	args := os.Args

	asciiArt.Input(args)
}
