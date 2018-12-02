package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

const inputFile = "input"

func main() {
	c, err := newCalibrator(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	c.calibrate()
}

type calibrator struct {
	rawData          []string
	processedData    []int64
	frequencyHistory map[int64]int
}

func newCalibrator(path string) (*calibrator, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var rawData []string
	var processedData []int64
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		rawData = append(rawData, scanner.Text())
		val, err := strconv.ParseInt(scanner.Text(), 10, 64)
		if err != nil {
			return nil, err
		}
		processedData = append(processedData, val)
	}
	return &calibrator{rawData: rawData, processedData: processedData}, scanner.Err()
}

func (c *calibrator) calibrate() {
	var frequency int64
	c.frequencyHistory = make(map[int64]int)

	for c.frequencyHistory[frequency] != 2 {
		for _, change := range c.processedData {
			currentFrequency := frequency
			frequency += change
			c.frequencyHistory[frequency]++
			fmt.Printf("Current frequency %v, change of %v; resulting frequency %v\n", currentFrequency, change, frequency)
			if c.frequencyHistory[frequency] == 2 {
				break
			}
		}
	}
	fmt.Printf("First duplicate frequency is %v\n", frequency)
}
