package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func readLines(inFile string) []string {
	var contents []string

	if file, err := os.Open(inFile); err != nil {
		log.Fatal(err)
	} else {
		defer file.Close()

		scanner := bufio.NewScanner(file)

		for scanner.Scan() {
			contents = append(contents, scanner.Text())
		}

		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
	}
	return contents
}

func day03() {
	fmt.Println("== DAY 03 ==")

	lines := readLines("./inputs/day-03-input.txt")
	lineCount := len(lines)
	columnCount := len(lines[0])

	var sums = make([]int, columnCount)

	for _, line := range lines {
		for pos, char := range line {
			if char == '1' {
				sums[pos]++
			}
		}
	}

	gamma := ""
	epsilon := ""

	for _, val := range sums {
		if val > lineCount/2 {
			gamma += "1"
			epsilon += "0"
		} else {
			gamma += "0"
			epsilon += "1"
		}
	}

	if gamma_val, err := strconv.ParseInt(gamma, 2, 32); err != nil {
		log.Fatal(err)
	} else {
		if epsilon_val, err := strconv.ParseInt(epsilon, 2, 32); err != nil {
			log.Fatal(err)
		} else {
			fmt.Printf("Gamma: %d, Epsilon: %d, Answer: %d\n", gamma_val, epsilon_val, gamma_val*epsilon_val)
		}
	}
}
