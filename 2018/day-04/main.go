package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"
)

const inputFile = "input"

type guard struct {
	guardID         int
	minutesSlept    float64
	sleepiestMinute int
}

type event struct {
	guardID      int
	start        time.Time
	end          time.Time
	minute       int
	sleepMinutes float64
	isAwake      bool
	eventStr     string
}

type events []event

func (e events) Len() int {
	return len(e)
}

func (e events) Less(i, j int) bool {
	return e[i].start.Before(e[j].start)
}

func (e events) Swap(i, j int) {
	e[i], e[j] = e[j], e[i]
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
	sorted := make(events, 0, len(data))
	for _, line := range data {
		e := parseEvent(line)
		sorted = append(sorted, e)
	}
	sort.Sort(sorted)

	// Must sort slice first to get missing Guard IDs and End Time
	parsed := make(events, 0, len(data))
	var eventGuardID int
	var guardID int
	for i, e := range sorted {
		parsedGuardID := parseGuardID(e.eventStr)
		if parsedGuardID == 0 {
			guardID = eventGuardID
		} else {
			guardID = parsedGuardID
			eventGuardID = parsedGuardID
		}
		e.guardID = guardID
		next := sorted[i:]
		if len(next) > 1 {
			e.end = sorted[i+1].start
		}
		if e.isAwake != true {
			e.sleepMinutes = e.end.Sub(e.start).Minutes()
		}
		parsed = append(parsed, e)
	}

	sleepy := getSleepiestGuard(parsed)
	fmt.Printf("The sleepiest guard is %v who slept for %v minutes\n", sleepy.guardID, sleepy.minutesSlept)
	fmt.Printf("The guard slept the most often during the %v minute\n", sleepy.sleepiestMinute)
	fmt.Printf("The answer is: %v\n", sleepy.guardID*sleepy.sleepiestMinute)

	sleepy2 := getSleepiestGuard2(parsed)
	fmt.Printf("The sleepiest guard is %v\n", sleepy2.guardID)
	fmt.Printf("The guard slept the most often during the %v minute\n", sleepy2.sleepiestMinute)
	fmt.Printf("The answer is: %v\n", sleepy2.guardID*sleepy2.sleepiestMinute)
}

func parseEvent(line string) event {
	valStr := strings.FieldsFunc(line, split)
	timeStr := valStr[0]
	eventStr := valStr[1]

	t, err := time.Parse("2006-01-02 15:04", timeStr)
	if err != nil {
		log.Fatal(err)
	}
	return event{
		start:    t,
		minute:   t.Minute(),
		isAwake:  strings.Contains(eventStr, "wakes") || strings.Contains(eventStr, "begins"),
		eventStr: eventStr,
	}
}

func parseGuardID(eventStr string) int {
	var guardID int

	re := regexp.MustCompile("[0-9]+")
	guardStr := re.FindString(eventStr)
	if guardStr != "" {
		guardID, err := strconv.Atoi(guardStr)
		if err != nil {
			log.Fatal(err)
		}
		return guardID
	}
	return guardID
}

func getSleepiestGuard(parsed events) guard {
	// find sleepiest guard
	guards := map[int]float64{}
	for _, e := range parsed {
		guards[e.guardID] += e.sleepMinutes
	}
	var sleepy guard
	for g, s := range guards {
		if s > sleepy.minutesSlept {
			sleepy.minutesSlept = s
			sleepy.guardID = g
		}
	}

	// filter to only sleepy events for our sleepy af guard
	var sleepyEvents events
	for _, e := range parsed {
		if sleepy.guardID == e.guardID {
			sleepyEvents = append(sleepyEvents, e)
		}
	}

	minutes := map[int]float64{}
	for _, e := range sleepyEvents {
		for d := e.start; d.Before(e.end) || d.Equal(e.end); d = d.Add(time.Minute * 1) {
			minutes[d.Minute()]++
		}
	}

	var highestMinutes float64
	for m, s := range minutes {
		if s > highestMinutes {
			highestMinutes = s
			sleepy.sleepiestMinute = m
		}
	}
	return sleepy
}

// Puzzle 2
func getSleepiestGuard2(parsed events) guard {
	// find sleepiest guard
	guards := map[int]map[int]int{}
	for _, e := range parsed {
		if guards[e.guardID] == nil {
			guards[e.guardID] = make(map[int]int)
		}
		if e.isAwake == false {
			for d := e.start; d.Before(e.end); d = d.Add(time.Minute * 1) {
				guards[e.guardID][d.Minute()]++
			}
		}
	}

	var sleepy guard
	var highestMinutes int
	for g, minutes := range guards {
		for m, s := range minutes {
			if s > highestMinutes {
				highestMinutes = s
				sleepy.guardID = g
				sleepy.sleepiestMinute = m
			}
		}
	}
	return sleepy
}

func split(r rune) bool {
	switch r {
	case '[', ']':
		return true
	}
	return false
}
