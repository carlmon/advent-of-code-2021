package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func loadOctoData(inFile string) ([][]byte, error) {
	file, err := os.Open(inFile)
	if err != nil {
		log.Fatal(err)
		return nil, err
	} else {
		defer file.Close()

		scanner := bufio.NewScanner(file)
		var data [][]byte

		for scanner.Scan() {
			var array = make([]byte, 10)
			for ix, chr := range scanner.Text() {
				array[ix] = (byte)(chr - '0')
			}
			data = append(data, array)
		}

		return data, scanner.Err()
	}
}

func flashRecurse(data [][]byte, pt Point) int {
	total := 1
	data[pt.x][pt.y] = 0

	for x := -1; x <= 1; x++ {
		for y := -1; y <= 1; y++ {
			x_coord := pt.x + x
			y_coord := pt.y + y
			if 0 <= x_coord && x_coord <= 9 && 0 <= y_coord && y_coord <= 9 {
				// cannot flash twice in one step
				if data[x_coord][y_coord] != 0 {
					data[x_coord][y_coord]++
				}
				if data[x_coord][y_coord] > 9 {
					total += flashRecurse(data, Point{x_coord, y_coord})
				}
			}
		}
	}

	return total
}

func runOctoStep(data [][]byte) int {
	for x := range data {
		for y := range data[x] {
			data[x][y]++
		}
	}

	flashes := 0
	for x := range data {
		for y := range data[x] {
			if data[x][y] > 9 {
				flashes += flashRecurse(data, Point{x, y})
			}
		}
	}
	return flashes
}

func runOctoSimulation(data [][]byte, steps int) {
	flashes := 0
	allFlashed := false

	for i := 0; i < steps; i++ {
		newFlashes := runOctoStep(data)
		if newFlashes == 100 {
			fmt.Printf("All flashed in step %d\n", i+1)
			allFlashed = true
		}
		flashes += newFlashes
	}
	fmt.Printf("Flashes in %d steps: %d\n", steps, flashes)

	step := 100
	for !allFlashed {
		newFlashes := runOctoStep(data)
		if newFlashes == 100 {
			fmt.Printf("All flashed in step %d\n", step+1)
			allFlashed = true
		}
		step++
	}
}

func init() {
	fmt.Println("== DAY 10 ==")

	if data, err := loadOctoData("./inputs/day-11-input.txt"); err != nil {
		log.Fatal(err)
	} else {
		runOctoSimulation(data, 100)
	}
}
