package main

import (
	"fmt"
	"bufio"
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
		var lineFirstInt int
		var lineLastInt int

		for i := 0; i <= len(line) - 1; i ++ {
			characterString := string(line[i])
			charInt, err := strconv.Atoi(characterString)
			if err != nil {
				continue
			}
			if lineFirstInt == 0 {
			    lineFirstInt = charInt
				lineLastInt = charInt
				continue
			}
			lineLastInt = charInt
		}
		// fmt.Println(lineFirstInt)
		// fmt.Println(lineLastInt)
		
		strA := strconv.Itoa(lineFirstInt)
		strB := strconv.Itoa(lineLastInt)
		lineString := strA + strB
		lineInt, err := strconv.Atoi(lineString)
		if err != nil {
			fmt.Println("Can't convert back to int")
		}

		// fmt.Println(lineInt)
		sum = sum + lineInt 
	}

	fmt.Println("Sum: ", sum)
}
