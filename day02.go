package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	gameRE := `(Game\s\d+)|(\d+\s\w+)`
	re := regexp.MustCompile(gameRE)
	part1 := 0
	part2 := 0
	var game, redCount, greenCount, blueCount int

	file, err := os.Open("day02.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		matches := re.FindAllString(scanner.Text(), -1)
		redCount, greenCount, blueCount = 0, 0, 0
		for _, v := range matches {
			sides := strings.Split(v, " ")
			if "Game" == sides[0] {
				game, err = strconv.Atoi(sides[1])
				if err != nil {
					panic(err)
				}
			} else {
				cubeCount, err := strconv.Atoi(sides[0])
				if err != nil {
					panic(err)
				}
				if "red" == sides[1] {
					redCount = max(redCount, cubeCount)
				} else if "green" == sides[1] {
					greenCount = max(greenCount, cubeCount)
				} else if "blue" == sides[1] {
					blueCount = max(blueCount, cubeCount)
				}
			}
		}
		if redCount <= 12 && greenCount <= 13 && blueCount <= 14 {
			part1 += game
		}
		part2 += redCount * greenCount * blueCount
	}
	fmt.Println(part1)
	fmt.Println(part2)
}
