package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func readFishData(inFile string) ([]uint64, error) {
	var data = make([]uint64, 9)

	file, err := os.Open(inFile)
	if err == nil {
		defer file.Close()

		scanner := bufio.NewScanner(file)

		scanner.Scan()
		for _, datum := range strings.Split(scanner.Text(), ",") {
			if val, err := strconv.Atoi(datum); err != nil {
				return data, err
			} else {
				data[val]++
			}
		}
	}

	return data, err
}

func runSimulation(fishData []uint64, days int) {
	// use a ring buffer for number of fish
	// each cycle adds the amount of currently breeding fish to the pool
	for i := 0; i < days; i++ {
		fishData[(i+7)%9] += fishData[i%9]
	}

	var sum uint64 = 0
	for _, val := range fishData {
		sum += val
	}

	fmt.Printf("Lanternfish after %d days: %d\n", days, sum)
}

func init() {
	fmt.Println("== DAY 06 ==")

	if fishData, err := readFishData("./inputs/day-06-input.txt"); err != nil {
		log.Fatal(err)
	} else {
		fishDataCopy := make([]uint64, 9)
		copy(fishDataCopy, fishData)

		runSimulation(fishData, 80)
		runSimulation(fishDataCopy, 256)
	}
}
