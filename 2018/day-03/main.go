package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const inputFile = "input"

type claim struct {
	id, x, y, w, h int
}

func main() {
	data, err := readInput(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	solve(data)
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
func solve(data []string) {
	var claims []claim

	for _, line := range data {
		c := parseClaim(line)
		claims = append(claims, c)
	}
	fabric := mapClaims(claims)
	uniqueClaim := findUniqueClaim(claims, fabric)
	fmt.Printf("The unqiue claim is %v\n", uniqueClaim)
}

func parseClaim(line string) claim {
	valStr := strings.FieldsFunc(line, split)
	val := stringsToInts(valStr)
	return claim{
		id: val[0],
		x:  val[1],
		y:  val[2],
		w:  val[3],
		h:  val[4],
	}
}

func mapClaims(claims []claim) [1000][1000]int {
	var fabric [1000][1000]int
	var overlap int

	for _, c := range claims {
		for x := c.x; x < c.x+c.w; x++ {
			for y := c.y; y < c.y+c.h; y++ {
				if fabric[x][y] == 0 {
					// first time we see a coordinate
					fabric[x][y] = 1
				} else if fabric[x][y] == 1 {
					// second time we see a coordinate
					fabric[x][y]++
					overlap++
				}
			}
		}
	}
	fmt.Printf("There are %v overlapping square inches\n", overlap)
	return fabric
}

func stringsToInts(values []string) []int {
	var ints []int
	for _, v := range values {
		i, err := strconv.Atoi(v)
		if err != nil {
			log.Fatal(err)
		}
		ints = append(ints, i)
	}
	return ints
}

func split(r rune) bool {
	switch r {
	case '#', '@', ',', ':', 'x', ' ':
		return true
	}
	return false
}

// Puzzle 2
func findUniqueClaim(claims []claim, fabric [1000][1000]int) int {
	var dupeCount int

	for _, c := range claims {
		for x := c.x; x < c.x+c.w; x++ {
			for y := c.y; y < c.y+c.h; y++ {
				if fabric[x][y] == 1 {
					continue
				} else if fabric[x][y] > 1 {
					dupeCount++
				}
			}
		}
		if dupeCount == 0 {
			return c.id
		}
		dupeCount = 0
	}
	return 0
}
