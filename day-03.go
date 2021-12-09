package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
)

func readLines(inFile string) ([]string, error) {
	var contents []string

	file, err := os.Open(inFile)
	if err == nil {
		defer file.Close()

		scanner := bufio.NewScanner(file)

		for scanner.Scan() {
			contents = append(contents, scanner.Text())
		}

		err = scanner.Err()
	}

	return contents, err
}

func mostCommon(lines []string, pos int) int {
	sum := 0
	lineCount := len(lines)

	for _, line := range lines {
		if line[pos] == '1' {
			sum++
		}
	}

	if sum*2 >= lineCount {
		return 1
	} else {
		return 0
	}
}

func partOne(lines []string) {
	columnCount := len(lines[0])
	gamma := 0
	epsilon := 0

	for i := 0; i < columnCount; i++ {
		gamma = gamma << 1
		epsilon = epsilon << 1
		if mostCommon(lines, i) == 1 {
			gamma++
		} else {
			epsilon++
		}
	}

	fmt.Printf("Gamma: %d, Epsilon: %d, Answer: %d\n", gamma, epsilon, gamma*epsilon)
}

func filterLines(lines []string, pos int, testChar byte) []string {
	var result = make([]string, 0)

	for _, line := range lines {
		if line[pos] == testChar {
			result = append(result, line)
		}
	}

	return result
}

func calcOxygenOrCo2(lines []string, flip bool) (int64, error) {
	columnCount := len(lines[0])
	for i := 0; i < columnCount; i++ {
		common := mostCommon(lines, i)
		// Flip to least common when searching for CO2 rating
		if flip {
			common = 1 - common
		}
		commonStr := strconv.Itoa(common)[0]
		lines = filterLines(lines, i, commonStr)
		if len(lines) == 1 {
			return strconv.ParseInt(lines[0], 2, 32)
		}
	}

	return 0, errors.New("rating value not found in lines")
}

func partTwo(lines []string) {

	if oxy, err := calcOxygenOrCo2(lines, false); err == nil {
		if co2, err := calcOxygenOrCo2(lines, true); err == nil {
			fmt.Printf("Oxygen Rating: %d, CO2 Rating: %d, Answer: %d\n", oxy, co2, oxy*co2)
		} else {
			log.Fatal(err)
		}
	} else {
		log.Fatal(err)
	}
}

func init() {
	fmt.Println("== DAY 03 ==")

	if lines, err := readLines("./inputs/day-03-input.txt"); err == nil {
		partOne(lines)
		partTwo(lines)
	} else {
		log.Fatal(err)
	}
}
