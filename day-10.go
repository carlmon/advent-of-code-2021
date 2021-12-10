package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

func errorScore(chr rune) int {
	switch chr {
	case ')':
		return 3
	case ']':
		return 57
	case '}':
		return 1197
	case '>':
		return 25137
	default:
		return 0
	}
}

func calculateLineError(line string) (int, []rune) {
	var stack = make([]rune, 0)
	for _, chr := range line {
		switch chr {
		case '[', '(', '<', '{':
			stack = append(stack, chr)
		case ']', ')', '>', '}':
			diff := chr - stack[len(stack)-1]
			if diff == 1 || diff == 2 {
				stack = stack[:len(stack)-1]
			} else {
				return errorScore(chr), stack
			}
		}
	}

	return 0, stack
}

func calculateRemainScore(remainder []rune) int {
	total := 0
	for i := len(remainder) - 1; i >= 0; i-- {
		total *= 5
		switch remainder[i] {
		case '(':
			total += 1
		case '[':
			total += 2
		case '{':
			total += 3
		case '<':
			total += 4
		}
	}
	return total
}

func calculateErrors(inFile string) {
	file, err := os.Open(inFile)
	if err != nil {
		log.Fatal(err)
	} else {
		defer file.Close()

		scanner := bufio.NewScanner(file)
		total := 0
		var completions = make([]int, 0)

		for scanner.Scan() {
			if score, remain := calculateLineError(scanner.Text()); score > 0 {
				total += score
			} else {
				completionScore := calculateRemainScore(remain)
				completions = append(completions, completionScore)
			}
		}

		fmt.Printf("Error score: %d\n", total)
		sort.Ints(completions)
		fmt.Printf("Middle score: %d\n", completions[len(completions)/2])
	}
}

func init() {
	fmt.Println("== DAY 10 ==")

	calculateErrors("./inputs/day-10-input.txt")
}
