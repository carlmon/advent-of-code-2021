package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	x int
	y int
}

type Line struct {
	a Point
	b Point
}

func readLinesData(inFile string) ([]Line, error) {
	var lines []Line

	file, err := os.Open(inFile)
	if err == nil {
		defer file.Close()
		scanner := bufio.NewScanner(file)

		for scanner.Scan() {
			if line, err := parseLine(scanner.Text()); err != nil {
				return lines, err
			} else {
				lines = append(lines, line)
			}
		}

		err = scanner.Err()
	}

	return lines, err
}

func parseLine(text string) (Line, error) {
	pointsStr := strings.Split(text, " -> ")
	if a, err := parsePoint(pointsStr[0]); err != nil {
		return Line{Point{0, 0}, Point{0, 0}}, err
	} else {
		if b, err := parsePoint(pointsStr[1]); err != nil {
			return Line{Point{0, 0}, Point{0, 0}}, err
		} else {
			if a.x < b.x {
				return Line{a, b}, nil
			} else if a.x == b.x {
				if a.y <= b.y {
					return Line{a, b}, nil
				} else {
					return Line{b, a}, nil
				}
			} else {
				return Line{b, a}, nil
			}
		}
	}
}

func parsePoint(pointComma string) (Point, error) {
	if x, err := strconv.Atoi(strings.Split(pointComma, ",")[0]); err != nil {
		return Point{0, 0}, err
	} else {
		if y, err := strconv.Atoi(strings.Split(pointComma, ",")[1]); err != nil {
			return Point{0, 0}, err
		} else {
			return Point{x, y}, nil
		}
	}
}

func calculateIntersections(lines []Line, point Point) (int, int) {
	noDiagonals := 0
	allOverlaps := 0

	for _, line := range lines {
		// ignore diagonals

		if (line.a.x <= point.x && point.x <= line.b.x) && ((line.a.y <= point.y && point.y <= line.b.y) || (line.a.y >= point.y && point.y >= line.b.y)) {
			if (point.y-line.a.y)*(line.b.x-line.a.x) == (point.x-line.a.x)*(line.b.y-line.a.y) {
				allOverlaps++

				if line.a.x == line.b.x || line.a.y == line.b.y {
					noDiagonals++
				}
			}
		}

		// no need to find more than two
		if noDiagonals > 1 {
			break
		}
	}

	return noDiagonals, allOverlaps
}

func doPartOne(lines []Line) {
	gridSize := 1000
	allOverlaps := 0
	straightOverlaps := 0
	for x := 0; x < gridSize; x++ {
		for y := 0; y < gridSize; y++ {
			if straightCount, diagonalCount := calculateIntersections(lines, Point{x, y}); straightCount > 1 {
				allOverlaps++
				straightOverlaps++
			} else if diagonalCount > 1 {
				allOverlaps++
			}
		}
	}

	fmt.Printf("Straight overlaps: %d, Plus diagonals: %d\n", straightOverlaps, allOverlaps)
}

func init() {
	fmt.Println("== DAY 05 ==")
	//if lines, err := readLinesData("./test.txt"); err != nil {
	if lines, err := readLinesData("./inputs/day-05-input.txt"); err != nil {
		log.Fatal(err)
	} else {
		doPartOne(lines)
	}
}
