package main

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"

	"github.com/jpillora/puzzler/harness/aoc"
)

func main() {
	aoc.Harness(run)
}

// on code change, run will be executed 4 times:
// 1. with: false (part1), and example input
// 2. with: true (part2), and example input
// 3. with: false (part1), and user input
// 4. with: true (part2), and user input
// the return value of each run is printed to stdout

func run(part2 bool, input string) any {
	score := 0

	if part2 {
		for _, line := range strings.Split(input, "\n") {
			if line == "" {
				continue
			}
			score += getScorePart2(line)
		}
		return score
	}

	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			return score
		}
		score += getScore(line)
	}

	// solve part 1 here
	return score
}

// 12 red cubes, 13 green cubes, and 14 blue cube
var red = 12
var green = 13
var blue = 14

func getScorePart2(line string) int {
	arr := []byte(line)
	split := bytes.Split(arr, []byte(":"))
	trims := bytes.Split((bytes.Trim(split[1], " ")), []byte(";"))

	var maxRed, maxGreen, maxBlue int

	for _, pull := range trims {
		for _, dice := range bytes.Split(pull, []byte(",")) {
			dice := bytes.Trim(dice, " ")
			colorGroup := bytes.Split(dice, []byte(" "))

			number, _ := strconv.Atoi(string(colorGroup[0]))
			color := colorGroup[1]

			switch string(color) {
			case "red":
				if number > maxRed {
					maxRed = number
				}

			case "green":
				if number > maxGreen {
					maxGreen = number
				}

			case "blue":
				if number > maxBlue {
					maxBlue = number
				}
			}

		}
	}

	return maxRed * maxGreen * maxBlue

}

func getScore(line string) int {
	arr := []byte(line)
	split := bytes.Split(arr, []byte(":"))
	id := bytes.Split(split[0], []byte(" "))[1]
	trims := bytes.Split((bytes.Trim(split[1], " ")), []byte(";"))

	intId, _ := strconv.Atoi(string(id))

	for _, pull := range trims {
		for _, dice := range bytes.Split(pull, []byte(",")) {
			dice := bytes.Trim(dice, " ")
			colorGroup := bytes.Split(dice, []byte(" "))

			number, _ := strconv.Atoi(string(colorGroup[0]))
			color := colorGroup[1]

			switch string(color) {
			case "red":
				if number > red {
					return 0
				}

			case "green":
				if number > green {
					return 0
				}

			case "blue":
				if number > blue {
					return 0
				}
			}

		}
	}

	return intId

}
