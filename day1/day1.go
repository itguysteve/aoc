package day1

import (
	"aoc/utils"
	"strings"
	"unicode"
)

func CheckForWrittenInteger(s string) int {
	digits := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}

	digit, found := digits[s]
	if found {
		return digit
	}
	// No digits are less than 3 letters, and i isn't a match, abort search
	if len(s) == 3 {
		return 0
	}
	// Remove the last character from the string and try again
	return CheckForWrittenInteger(s[0 : len(s)-1])
}

func Solve() (int, int) {
	input := utils.GetInputData(1)
	instructions := strings.Split(input, "\n")
	intTotal := 0
	strTotal := 0

	for _, instruction := range instructions {
		// Maps to hold first and last found integers from instructions
		intInts := map[int]int{0: 0, 1: 0}
		strInts := map[int]int{0: 0, 1: 0}
		var number int

		for index, r := range instruction {
			if unicode.IsDigit(r) {
				number = int(r - '0')
				if intInts[0] == 0 {
					intInts[0] = number
					intInts[1] = number
				} else {
					intInts[1] = number
				}
			} else { // Not a digit, search for written digits
				var test string

				tail := len(instruction) - index
				if tail > 4 {
					test = instruction[index : index+5]
				} else if tail > 2 {
					// Don't have 5 characters available to send, reducing search string length
					test = instruction[index : index+tail]
				} else {
					// Remainder of this string is too short to contain a written character
					continue
				}
				number = CheckForWrittenInteger(test)
				if number == 0 {
					// Did not find a written digit
					continue
				}
			}

			if strInts[0] == 0 {
				strInts[0] = number
				strInts[1] = number
			} else {
				strInts[1] = number
			}
		}
		intSum := intInts[0]*10 + intInts[1]
		intTotal += intSum

		strSum := strInts[0]*10 + strInts[1]
		strTotal += strSum
	}
	return intTotal, strTotal
}
