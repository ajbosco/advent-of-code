package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const inputFile = "input"

func main() {
	data, err := readInput(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	checkSum(data)
	findMatch(data)
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

func checkSum(data []string) {
	var twoCount int
	var threeCount int

	for _, l := range data {
		characterMap := countCharacters(strings.Split(l, ""))
		twoChar, threeChar := lineCheck(characterMap)
		if twoChar {
			twoCount++
		}
		if threeChar {
			threeCount++
		}
	}
	checkSum := twoCount * threeCount
	fmt.Printf("The checksum is %v\n", checkSum)
}

// Puzzle 1
func countCharacters(characters []string) map[string]int {
	characterMap := make(map[string]int)
	for _, c := range characters {
		characterMap[c]++
	}
	return characterMap
}

func lineCheck(characterMap map[string]int) (bool, bool) {
	var twoChar bool
	var threeChar bool

	for _, v := range characterMap {
		if v == 2 {
			twoChar = true
		}
		if v == 3 {
			threeChar = true
		}
	}
	return twoChar, threeChar
}

// Puzzle 2
func findMatch(data []string) {
	for _, currentID := range data {
		for _, compareID := range data {
			match := compare(currentID, compareID)
			if match {
				fmt.Printf("Matching IDs: %q, %q\n", currentID, compareID)
			}
		}
	}
}

func compare(id1 string, id2 string) bool {
	id1Chars := strings.Split(id1, "")
	id2Chars := strings.Split(id2, "")
	diff := 0
	for i, c := range id1Chars {
		if c != id2Chars[i] {
			diff++
		}
	}
	if diff == 1 {
		return true
	}
	return false
}
