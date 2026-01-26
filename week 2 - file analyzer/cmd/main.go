package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

// A file opens a cursor at the start of the file.
// A scanner runs the scanner over the bytes in the file and can only run once (needs to be rewinded)
// Package 'io' defines reader and writer interfaces, which are implemented by other packages like: 'os', 'gzip', 'tar', etc.
// Package 'bufio' provides helper functions and structs for reading/writing. It is a layer between the OS and the application which buffers data. Like filling a glas with water instead of drinking from the sink directly.

func main() {
	fileName := flag.String("f", "", "File to analyse")
	lines := flag.Bool("l", false, "Count Lines")
	words := flag.Bool("w", false, "Count Words")
	chars := flag.Bool("c", false, "Count Chars")
	flag.Parse()

	f, err := os.Open(*fileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	if *lines {
		fmt.Println("Lines counted", countLines(f))
		f.Seek(0, io.SeekStart)
	}

	if *words {
		fmt.Println("Words counted", countWords(f))
		f.Seek(0, io.SeekStart)
	}

	if *chars {
		fmt.Println("Chars counted", countChars(f))
	}
}

func countLines(r io.Reader) int {
	lineCount := 0

	lineScanner := bufio.NewScanner(r)
	lineScanner.Split(bufio.ScanLines)
	for lineScanner.Scan() {
		lineCount++
	}
	if err := lineScanner.Err(); err != nil {
		fmt.Printf("Error scanning file: %v", err)
	}
	return lineCount
}

func countWords(r io.Reader) int {
	wordCount := 0

	wordScanner := bufio.NewScanner(r)
	wordScanner.Split(bufio.ScanWords)
	for wordScanner.Scan() {
		wordCount++
	}
	if err := wordScanner.Err(); err != nil {
		fmt.Printf("Error scanning file: %v", err)
	}
	return wordCount
}

func countChars(r io.Reader) int {
	charCount := 0

	charScanner := bufio.NewScanner(r)
	charScanner.Split(bufio.ScanRunes)
	for charScanner.Scan() {
		charCount++
	}
	if err := charScanner.Err(); err != nil {
		fmt.Printf("Error scanning file: %v", err)
	}
	return charCount
}
