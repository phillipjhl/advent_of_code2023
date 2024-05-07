package digitfinder

import (
	"fmt"
	"strconv"
	"strings"
)

var digits = [9]string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

type DigitsInLine struct {
	First      int
	Last       int
	FirstIndex int
	LastIndex  int
}

func FindDigitsInLine(line string) (*DigitsInLine, error) {
	foundDigits := &DigitsInLine{}

	// Find the digit that exists First in the line
	for i := range digits {
		// the temp index is the First index of the string where a match has occurred
		firstIndex := strings.Index(line, digits[i])
		lastIndex := strings.LastIndex(line, digits[i])
		// pass, digit not found
		if firstIndex == -1 || lastIndex == -1 {
			continue
		}

		if foundDigits.First == 0 || foundDigits.Last == 0 {

			// set the found digit only if it hasn't been set yet
			if foundDigits.First == 0 && foundDigits.FirstIndex == 0 {
				foundDigits.First = i + 1
				foundDigits.FirstIndex = firstIndex
			}

			// set the last found digit if it hasn't been set yet
			if foundDigits.Last == 0 && foundDigits.LastIndex == 0 {
				foundDigits.Last = i + 1
				foundDigits.LastIndex = firstIndex
			}

		}

		// the found digit is earlier in the list, set it as the output index
		if firstIndex < foundDigits.FirstIndex {
			foundDigits.FirstIndex = firstIndex
			// set a found digit
			foundDigits.First = i + 1
		}

		// the found digit is later in the string, set it as the output last index
		if lastIndex > foundDigits.LastIndex {
			foundDigits.LastIndex = lastIndex
			foundDigits.Last = i + 1
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
			lastIntIndex = i
			continue
		}
		lastInt = charInt
		lastIntIndex = i
	}

	fmt.Printf("First Int: %v \n", firstInt)
	fmt.Printf("Last Int: %v \n", lastInt)
	return firstInt, firstIntIndex, lastInt, lastIntIndex
}
