package digitfinder

import (
	"fmt"
	"strconv"
	"strings"
)

var digits = [9]string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

type DigitsInLine struct {
	first int
	last int
	firstIndex int
	lastIndex int
}

func FindDigitsInLine(line string) (DigitsInLine, error) {	
	var foundDigits DigitsInLine

	// Find the digit that exists first in the line
	for i := range digits {
		// the temp index is the first index of the string where a match has occurred
		tempIndex := strings.Index(line, digits[i])
		// pass, digit not found
		if tempIndex == -1 {
			continue
		}

		// set the found digit only if it hasn't been set yet
		if foundDigits.first == 0 && foundDigits.firstIndex == 0 {
			foundDigits.first= i + 1
			foundDigits.firstIndex = tempIndex
			foundDigits.last = i + 1
			foundDigits.lastIndex = tempIndex
			continue
		}

		// the found digit is earlier in the list, set it as the output index
		if tempIndex < foundDigits.firstIndex {
			foundDigits.firstIndex = tempIndex
			// set a found digit
			foundDigits.first = i + 1
		}

		// the found digit is later in the string, set it as the output last index
		if tempIndex > foundDigits.lastIndex {
			foundDigits.lastIndex = tempIndex
			foundDigits.last = i + 1
		}
	}

	return foundDigits, nil
}

func FindFirstAndLastInts(line string) (int, int, int, int) {
	var firstInt int
	var firstIntIndex int
	var lastInt int
	var lastIntIndex int

	for i := 0; i <= len(line)-1; i++ {
		characterString := string(line[i])
		charInt, err := strconv.Atoi(characterString)
		if err != nil {
			continue
		}
		if firstInt == 0 {
			firstInt = charInt
			firstIntIndex = i
			lastInt = charInt
			continue
		}
		lastInt = charInt
		lastIntIndex = i
	}

	fmt.Println(firstInt)
	fmt.Println(lastInt)
	return firstInt, firstIntIndex, lastInt, lastIntIndex
}
