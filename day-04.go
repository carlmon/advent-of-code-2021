package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func readBingoData(inFile string) ([]string, [][][]string, error) {
	var numbers []string
	var boards [][][]string
	var line string

	file, err := os.Open(inFile)
	if err == nil {
		defer file.Close()

		scanner := bufio.NewScanner(file)

		scanner.Scan()
		numbers = strings.Split(scanner.Text(), ",")

		boardIndex := -1
		for scanner.Scan() {
			line = scanner.Text()
			if len(line) == 0 {
				boardIndex++
				boards = append(boards, make([][]string, 0))
			} else {
				var rowNumbers []string
				for _, elem := range strings.Split(line, " ") {
					if elem != "" {
						rowNumbers = append(rowNumbers, elem)
					}
				}
				boards[boardIndex] = append(boards[boardIndex], rowNumbers)
			}
		}

		err = scanner.Err()
	}

	return numbers, boards, err
}

func runBingo(numbers []string, boards [][][]string) {
	numbersDict := make(map[string]int)
	for ix, number := range numbers {
		numbersDict[number] = ix
	}

	minIndex := 100
	maxIndex := 0
	firstWinningBoard := -1
	lastWinningBoard := -1

	for boardIndex, board := range boards {
		if bingo, index := checkBoard(numbersDict, board); bingo {
			if index < minIndex {
				minIndex = index
				firstWinningBoard = boardIndex
			} else if index > maxIndex {
				maxIndex = index
				lastWinningBoard = boardIndex
			}
		}
	}

	printResult(numbersDict, numbers, boards, firstWinningBoard, minIndex)
	printResult(numbersDict, numbers, boards, lastWinningBoard, maxIndex)
}

func checkBoard(numbersDict map[string]int, board [][]string) (bool, int) {
	minIndex := 100

	for i := 0; i < len(board); i++ {
		if bingo, index := checkRow(numbersDict, board, i); bingo {
			if index < minIndex {
				minIndex = index
			}
		}

		if bingo, index := checkCol(numbersDict, board, i); bingo {
			if index < minIndex {
				minIndex = index
			}
		}
	}

	return minIndex < 100, minIndex
}

func checkRow(numbersDict map[string]int, board [][]string, row int) (bool, int) {
	maxIndex := 0

	for i := 0; i < len(board); i++ {
		if val, prs := numbersDict[board[row][i]]; !prs {
			return false, 0
		} else if maxIndex < val {
			maxIndex = val
		}
	}

	return true, maxIndex
}

func checkCol(numbersDict map[string]int, board [][]string, col int) (bool, int) {
	maxIndex := 0

	for i := 0; i < len(board); i++ {
		if val, prs := numbersDict[board[i][col]]; !prs {
			return false, 0
		} else if maxIndex < val {
			maxIndex = val
		}
	}

	return true, maxIndex
}

func printResult(numbersDict map[string]int, numbers []string, boards [][][]string, boardIndex int, lastDrawIndex int) {
	if sum, err := sumUnmarked(numbersDict, boards[boardIndex], lastDrawIndex); err != nil {
		log.Fatal(err)
	} else {
		if winNumber, err := strconv.Atoi(numbers[lastDrawIndex]); err != nil {
			log.Fatal(err)
		} else {
			log.Printf("Board index: %d, Numbers drawn: %d, Answer: %d", boardIndex, lastDrawIndex, sum*winNumber)
		}
	}
}

func sumUnmarked(numbersDict map[string]int, board [][]string, maxIndex int) (int, error) {
	sum := 0

	for _, row := range board {
		for _, value := range row {
			if val, prs := numbersDict[value]; !prs || val > maxIndex {
				if valueInt, err := strconv.Atoi(value); err != nil {
					return 0, err
				} else {
					sum += valueInt
				}
			}
		}
	}

	return sum, nil
}

func init() {
	fmt.Println("== DAY 04 ==")

	if numbers, boards, err := readBingoData("./inputs/day-04-input.txt"); err != nil {
		log.Fatal(err)
	} else {
		runBingo(numbers, boards)
	}
}
