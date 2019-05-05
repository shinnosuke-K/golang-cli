package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

const otherWord = "*"

var transforms = []string{
	otherWord,
	otherWord,
	otherWord,
	otherWord,
	otherWord + "app",
	otherWord + "site",
	otherWord + "time",
	"get" + otherWord,
	"go" + otherWord,
	"lets" + otherWord,
}

func readFromFile(filePath string) []string {
	f, err := os.Open(filePath)
	if err != nil {
		fmt.Println("error")
		os.Exit(1)
	}
	defer f.Close()

	lines := make([]string, 0)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "File %s scan error: %v\n", filePath, err)
	}
	return lines
}

func main() {
	lines := readFromFile("text")
	rand.Seed(time.Now().UTC().UnixNano())
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		t := lines[rand.Intn(len(lines))]
		fmt.Println(strings.Replace(t, otherWord, s.Text(), -1))
	}
}
