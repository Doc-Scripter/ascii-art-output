package main

import (
	"bytes"
	"os"
	"testing"

	a "ascii/functions"
)

var testCases = []struct {
	name          string
	expectedLines string
}{
	{"hello", read()},
}

func TestParagraph(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			old := os.Stdout
			r, w, _ := os.Pipe()
			os.Stdout = w
			s, _ := os.ReadFile("resources/standard.txt")

			m := a.AsciiArt(string(s))
			a.Art("hello", m)
			w.Close()
			os.Stdout = old
			var got bytes.Buffer
			_, err := got.ReadFrom(r)
			if err != nil {
				t.Fatal("error reading from os out")
			}
			expected := tc.expectedLines
			if expected != got.String() {
				t.Errorf("expected\n%v\n got\n%v", expected, got.String())
			}
		})
	}
}

func read() string {
	word, err := os.ReadFile("resources/test.txt")
	if err != nil {
		panic("error reading file")
	}
	return string(word)
}

var testCases1 = []struct {
	name          string
	expectedLines string
}{
	{"hello there", read1()},
}

func TestParagraph1(t *testing.T) {
	for _, tc := range testCases1 {
		t.Run(tc.name, func(t *testing.T) {
			old := os.Stdout
			r, w, _ := os.Pipe()
			os.Stdout = w
			s, _ := os.ReadFile("resources/shadow.txt")

			m := a.AsciiArt(string(s))
			a.Art("hello there", m)
			w.Close()
			os.Stdout = old
			var got bytes.Buffer
			_, err := got.ReadFrom(r)
			if err != nil {
				t.Fatal("error reading from os out")
			}
			expected := tc.expectedLines
			if expected != got.String() {
				t.Errorf("expected\n%v\n got\n%v", expected, got.String())
			}
		})
	}
}

func read1() string {
	word, err := os.ReadFile("resources/test2.txt")
	if err != nil {
		panic("error reading file")
	}
	return string(word)
}

var testCases2 = []struct {
	name          string
	expectedLines string
}{
	{"ABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890abcdefghijklmnopqrstuvwxyz!@#$%^&*()_+=<>?", read2()},
}

func TestParagraph2(t *testing.T) {
	for _, tc := range testCases2 {
		t.Run(tc.name, func(t *testing.T) {
			old := os.Stdout
			r, w, _ := os.Pipe()
			os.Stdout = w
			s, _ := os.ReadFile("resources/thinkertoy.txt")

			m := a.AsciiArt(string(s))
			a.Art("ABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890abcdefghijklmnopqrstuvwxyz!@#$%^&*()_+=<>?", m)
			w.Close()
			os.Stdout = old
			var got bytes.Buffer
			_, err := got.ReadFrom(r)
			if err != nil {
				t.Fatal("error reading from os out")
			}
			expected := tc.expectedLines
			if expected != got.String() {
				t.Errorf("expected\n%v\n got\n%v", expected, got.String())
			}
		})
	}
}

func read2() string {
	word, err := os.ReadFile("resources/test3.txt")
	if err != nil {
		panic("error reading file")
	}
	return string(word)
}

var testCases3 = []struct {
	name          string
	expectedLines string
}{
	{"hello", read3()},
}

func TestParagraph3(t *testing.T) {
	for _, tc := range testCases3 {
		t.Run(tc.name, func(t *testing.T) {
			old := os.Stdout
			r, w, _ := os.Pipe()
			os.Stdout = w
			s, _ := os.ReadFile("resources/ac.txt")

			m := a.AsciiArt(string(s))
			a.Art("hello", m)
			w.Close()
			os.Stdout = old
			var got bytes.Buffer
			_, err := got.ReadFrom(r)
			if err != nil {
				t.Fatal("error reading from os out")
			}
			expected := tc.expectedLines
			if expected != got.String() {
				t.Errorf("expected\n%v\n got\n%v", expected, got.String())
			}
		})
	}
}

func read3() string {
	word, err := os.ReadFile("resources/test4.txt")
	if err != nil {
		panic("error reading file")
	}
	return string(word)
}
