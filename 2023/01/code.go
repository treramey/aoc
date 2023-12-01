package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"

	"github.com/jpillora/puzzler/harness/aoc"
	"github.com/spf13/cast"
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

func matcher(input string) (int, bool) {
	digits := []string{}
	for i := 0; i < len(input); i++ {
		charStr := string(input[i])
		_, err := strconv.Atoi(charStr)
		if err == nil {
			digits = append(digits, charStr)
		}
	}
	if len(digits) == 0 {
		return 0, false
	}

	result, err := strconv.Atoi(digits[0] + digits[len(digits)-1])

	return result, err == nil
}

func run(part2 bool, input string) any {
	lines := strings.Split(input, "\n")
	santaCode := 0

	if part2 {
		for _, line := range lines {
			digits := []string{}
			for i, char := range line {
				if unicode.IsDigit(char) && cast.ToInt(string(char)) != 0 && cast.ToInt(string(char)) < 10 {
					digits = append(digits, string(char))
				}
				for j, value := range []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"} {
					if strings.HasPrefix(line[i:], value) {
						digits = append(digits, fmt.Sprint(j+1))
					}
				}
			}
			if len(digits) > 0 {
				combined := string(digits[0]) + string(digits[len(digits)-1])
				score := cast.ToInt(combined)
				santaCode += score
			}
		}
		return santaCode
	}

	for _, line := range lines {
		number, match := matcher(line)

		if match {
			santaCode += number
		}

	}
	return santaCode
}
