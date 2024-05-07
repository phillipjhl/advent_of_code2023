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
		var digitsFoundInLine = make([]int, 2)

		// loops over each line to the first and last integer characters
		firstInt, firstIntIndex, lastInt, lastIntIndex := digitfinder.FindFirstAndLastInts(line)

		// finds the first and last word representation of digits
		foundDigits, err := digitfinder.FindDigitsInLine(line)

		// int was not found
		if firstInt == 0 && foundDigits.First == 0 {
			// do nothing
		} else if firstInt != 0 && foundDigits.First == 0 {
		// use int as it's the only option
			digitsFoundInLine[0] = firstInt
		} else if foundDigits.First != 0 && firstInt == 0 {
		// use digits as it's the only option
			digitsFoundInLine[0] = foundDigits.First
		}  else {
			if firstIntIndex > foundDigits.FirstIndex {
				digitsFoundInLine[0] = foundDigits.First
			} else {
				// strA := strconv.Itoa(firstInt)
				digitsFoundInLine[0] = firstInt
			}
		}
		// neither int or digit was not found
		if lastInt == 0 && foundDigits.Last == 0 {
			// set last digit
			digitsFoundInLine[1] = digitsFoundInLine[0]
		} else if lastInt != 0 && foundDigits.Last == 0 {
		// use int as it's the only option
			digitsFoundInLine[1] = lastInt
		} else if foundDigits.Last != 0 && lastInt == 0 {
		// use digits as it's the only option
			digitsFoundInLine[1] = foundDigits.Last
		} else {
			fmt.Printf("dg: %v \n", foundDigits)
			fmt.Printf("li: %v \n", lastIntIndex)
			if lastIntIndex > foundDigits.LastIndex {
				digitsFoundInLine[1] = lastInt
			} else {
				// strB := strconv.Itoa(lastInt)
				digitsFoundInLine[1] = foundDigits.Last
			}
		}

		// conjoin the two digits to get the final result of the line
		fmt.Printf("DigitsSlice: %v \n", digitsFoundInLine)
		var lineString string
		lineString = strconv.Itoa(digitsFoundInLine[0]) + strconv.Itoa(digitsFoundInLine[1])
		fmt.Printf("LineAsString: %s \n", lineString)
		lineInt, err := strconv.Atoi(lineString)
		if err != nil {
			fmt.Printf("Can't convert %v back to int \n", lineString)
		}

		fmt.Printf("Line Int: %v \n ", lineInt)
		sum = sum + lineInt

		fmt.Println("---")
	}

	fmt.Println("Sum: ", sum)
}

func FindFirstAndLastInts(line string) {
	panic("unimplemented")
}
