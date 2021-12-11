package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
)

func identify25(idMap map[int]string, word string) (int, error) {
	if four, rec := idMap[4]; rec {
		contains := 0
		for _, chr := range four {
			if strings.ContainsRune(word, chr) {
				contains++
			}
		}
		switch contains {
		case 2:
			return 2, nil
		case 3:
			return 5, nil
		}
	}
	return 0, errors.New("could not id 2/5")
}

func identify235(idMap map[int]string, word string) (int, error) {
	if seven, rec := idMap[7]; rec {
		contains := 0
		for _, chr := range seven {
			if strings.ContainsRune(word, chr) {
				contains++
			}
		}
		switch contains {
		case 3:
			return 3, nil
		case 2:
			return identify25(idMap, word)
		}
	}
	return 0, errors.New("could not id 2/3/5")
}

func identify09(idMap map[int]string, word string) (int, error) {
	if four, rec := idMap[4]; rec {
		contains := 0
		for _, chr := range four {
			if strings.ContainsRune(word, chr) {
				contains++
			}
		}
		switch contains {
		case 3:
			return 0, nil
		case 4:
			return 9, nil
		}
	}
	return 0, errors.New("could not id 2/5")
}

func identify069(idMap map[int]string, word string) (int, error) {
	if seven, rec := idMap[7]; rec {
		contains := 0
		for _, chr := range seven {
			if strings.ContainsRune(word, chr) {
				contains++
			}
		}
		switch contains {
		case 2:
			return 6, nil
		case 3:
			return identify09(idMap, word)
		}
	}
	return 0, errors.New("could not id 2/3/5")
}

func addToIdMap(idMap map[int]string, line string) {
	for _, word := range strings.Split(line, " ") {
		switch len(word) {
		case 2:
			idMap[1] = word
		case 3:
			idMap[7] = word
		case 4:
			idMap[4] = word
		case 7:
			idMap[8] = word
		}
	}
}

func solveRandomDigits(inFile string) {

	if file, err := os.Open(inFile); err != nil {
		log.Fatal(err)
	} else {
		defer file.Close()

		scanner := bufio.NewScanner(file)
		easyDigits := 0
		total := 0

		for scanner.Scan() {
			idMap := make(map[int]string)
			addToIdMap(idMap, scanner.Text())

			realDigits := 0
			segment2 := strings.Split(strings.Split(scanner.Text(), " | ")[1], " ")

			for _, word := range segment2 {
				realDigits *= 10
				switch len(word) {
				case 2:
					easyDigits++
					realDigits += 1
				case 3:
					easyDigits++
					realDigits += 7
				case 4:
					easyDigits++
					realDigits += 4
				case 7:
					easyDigits++
					realDigits += 8
				case 5:
					if val, err := identify235(idMap, word); err != nil {
						log.Fatal(err)
						return
					} else {
						realDigits += val
					}
				case 6:
					if val, err := identify069(idMap, word); err != nil {
						log.Fatal(err)
						return
					} else {
						realDigits += val
					}
				}
			}

			total += realDigits
		}

		if err = scanner.Err(); err != nil {
			log.Fatal(err)
		} else {
			fmt.Printf("Number of easy digits: %d\n", easyDigits)
			fmt.Printf("Sum of output: %d\n", total)
		}
	}

}

func init() {
	fmt.Println("== DAY 08 ==")
	solveRandomDigits("./inputs/day-08-input.txt")
}
