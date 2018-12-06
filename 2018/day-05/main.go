package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"
)

const (
	inputFile = "input"
	alphabet  = "abcdefghijklmnopqrstuvwxyz"
)

func main() {
	data, err := readInput(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	solve1(data)
	solve2(data)
}

func readInput(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var data []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		data = append(data, scanner.Text())
	}
	return data, err
}

// Puzzle 1
func solve1(data []string) {
	for _, l := range data {
		cleanStr := check(l)
		fmt.Printf("The answer is: %v\n", len(cleanStr))
	}
}

func check(s string) string {
	b := []byte(s)
CheckLoop:
	for {
		for i := 0; i < len(b)-1; i++ {
			if match(b[i], b[i+1]) {
				b = append(b[:i], b[i+2:]...)
				continue CheckLoop
			}
		}
		return string(b)
	}

}

func match(val1 byte, val2 byte) bool {
	r1 := rune(val1)
	r2 := rune(val2)
	if unicode.IsLower(r1) && unicode.IsUpper(r2) {
		if unicode.ToUpper(r1) == unicode.ToUpper(r2) {
			return true
		}
	}
	if unicode.IsUpper(r1) && unicode.IsLower(r2) {
		if unicode.ToUpper(r1) == unicode.ToUpper(r2) {
			return true
		}
	}
	return false
}

// Puzzle 2
func solve2(data []string) {
	for _, l := range data {
		winningScore := len(l)
		// we just gonna brute force this thing
		for _, letter := range alphabet {
			s1 := strings.Replace(l, string(letter), "", -1)
			s2 := strings.Replace(s1, string(unicode.ToUpper(letter)), "", -1)
			cleanStr := check(s2)
			curScore := len(cleanStr)
			if curScore < winningScore {
				winningScore = curScore
			}
		}
		fmt.Printf("The answer is: %v\n", winningScore)
	}
}
