package main

import (
	"os"

	asciiArt "ascii/functions"
)

func main() {
	args := os.Args

	asciiArt.Input(args)
}
