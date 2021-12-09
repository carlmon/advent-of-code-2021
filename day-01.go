package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

func countIncreases(inFile string) {
	file, err := os.Open(inFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	increases := 0
	last := math.MaxInt32

	for scanner.Scan() {
		current, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}

		if last < current {
			increases++
		}
		last = current
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Increases: %d\n", increases)
}

func createSlidingWindow(inFile, outFile string) {
	file, err := os.Open(inFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	windowFile, err := os.OpenFile(outFile, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}

	datawriter := bufio.NewWriter(windowFile)
	defer windowFile.Close()

	index := 0
	windows := [3]int{0, 0, 0}

	for scanner.Scan() {
		val, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}

		if 1 <= index {
			windows[(index-1)%3] += val
		}
		if 2 <= index {
			windows[(index-2)%3] += val
		}

		datawriter.WriteString(strconv.Itoa(windows[index%3]) + "\n")

		windows[index%3] = val
		index++
	}
	datawriter.Flush()

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func init() {
	fmt.Println("== DAY 01 ==")
	countIncreases("./inputs/day-01-input.txt")
	createSlidingWindow("./inputs/day-01-input.txt", "./day-01-sliding.txt")
	countIncreases("./day-01-sliding.txt")
}
