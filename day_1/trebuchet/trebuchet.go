package main

import (
	"bufio"
	"fmt"
	"github.com/phillipjhl/digitfinder"
	"os"
	"strconv"
)

func main() {
	// read each line of the input to find the sum of all calculated lines

	file, err := os.Open("input.txt")

	if err != nil {
		fmt.Println("Error opening file")
		return
	}

	defer file.Close()

	// process each line of the file by reading each bufio
	scanner := bufio.NewScanner(file)

	sum := 0

	for scanner.Scan() {
		line := scanner.Text()

		fmt.Println(line)

		// slice to hold the first and last found digits in the line
		var digitsFoundInLine = make([]string, 2)

		// loops over each line to the first and last integer characters
		firstInt, firstIntIndex, lastInt, lastIntIndex := digitfinder.FindFirstAndLastInts(line)

		// finds the first and last word representation of digits
		foundDigits, err := digitfinder.FindDigitsInLine(line)

		// int was not found
		if firstIntIndex == 0 {
			// set first digit
			digitsFoundInLine[0] = foundDigits.first
		} else {
			if firstIntIndex > foundDigits.firstIndex {
				digitsFoundInLine[0] = foundDigits.first
			} else {
				strA := strconv.Itoa(firstInt)
				digitsFoundInLine[0] = strA
			}
		}
		// int was not found
		if lastIntIndex == 0 {
			// set last digit
			digitsFoundInLine[1] = foundDigits.last
		} else {
			if lastIntIndex > foundDigits.lastIndex {
				digitsFoundInLine[1] = foundDigits.last
			} else {
				strB := strconv.Itoa(lastInt)
				digitsFoundInLine[1] = strB
			}
		}

		// conjoin the two digits to get the final result of the line
		fmt.Printf("DigitsSlice: %v \n", digitsFoundInLine)
		lineString := digitsFoundInLine[0] + digitsFoundInLine[1]
		fmt.Printf("LineAsString: %s \n", lineString)
		lineInt, err := strconv.Atoi(lineString)
		if err != nil {
			fmt.Printf("Can't convert %v back to int \n", lineString)
		}

		fmt.Printf("Line Int: %v \n ", lineInt)
		sum = sum + lineInt
	}

	fmt.Println("Sum: ", sum)
}

func FindFirstAndLastInts(line string) {
	panic("unimplemented")
}
