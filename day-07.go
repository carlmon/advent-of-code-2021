package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func readCrabPositions(inFile string) (map[int]uint64, error) {
	var data = make(map[int]uint64)

	file, err := os.Open(inFile)
	if err == nil {
		defer file.Close()

		scanner := bufio.NewScanner(file)

		scanner.Scan()
		for _, position := range strings.Split(scanner.Text(), ",") {
			if val, err := strconv.Atoi(position); err != nil {
				return data, err
			} else {
				data[val]++
			}
		}
	}

	return data, err
}

func fixedMovementCost(n int) uint64 {
	return (uint64)(n * (n + 1) / 2)
}

func calculateCost(positions map[int]uint64, target int) (uint64, uint64) {
	var cost uint64 = 0
	var cost_fixed uint64 = 0
	for key, val := range positions {
		if target > key {
			cost += (uint64)(target-key) * val
			cost_fixed += fixedMovementCost(target-key) * val
		} else {
			cost += (uint64)(key-target) * val
			cost_fixed += fixedMovementCost(key-target) * val
		}
	}
	return cost, cost_fixed
}

func findOptimalMove(positions map[int]uint64) {
	var lastCost uint64 = math.MaxInt32
	index := 0
	for {
		cost, _ := calculateCost(positions, index)
		if lastCost < cost {
			fmt.Printf("Position: %d, Cost: %d\n", index-1, lastCost)
			break
		} else {
			lastCost = cost
			index++
		}
	}

	lastCost = math.MaxInt32
	index = 0
	for {
		_, cost := calculateCost(positions, index)
		if lastCost < cost {
			fmt.Printf("Position: %d, Fixed cost: %d\n", index-1, lastCost)
			break
		} else {
			lastCost = cost
			index++
		}
	}
}

func init() {
	fmt.Println("== DAY 07 ==")
	if positions, err := readCrabPositions("./inputs/day-07-input.txt"); err != nil {
		log.Fatal(err)
	} else {
		findOptimalMove(positions)
	}
}
