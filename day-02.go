package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func init() {
	fmt.Println("== DAY 02 ==")

	file, err := os.Open("./inputs/day-02-input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	horiz := 0
	depth := 0
	depth_fixed := 0
	aim := 0

	for scanner.Scan() {
		movement := strings.Split(scanner.Text(), " ")
		amount, err := strconv.Atoi(movement[1])
		if err != nil {
			log.Fatal(err)
		}

		switch movement[0] {
		case "down":
			depth += amount
			aim += amount
		case "up":
			depth -= amount
			aim -= amount
		case "forward":
			horiz += amount
			depth_fixed += aim * amount
		case "back":
			horiz -= amount
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Depth: %d, Horiz: %d, Answer: %d\n", depth, horiz, depth*horiz)
	fmt.Printf("Aim: %d, Depth: %d, Horiz: %d, Answer: %d\n", aim, depth_fixed, horiz, depth_fixed*horiz)
}
