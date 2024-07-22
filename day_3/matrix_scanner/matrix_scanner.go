package main

import (
	"fmt"
	"os"
    "bufio"
	// "strings"
)

func main() {
	// Open the file
	filePath := os.Args[1]
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error: %v", err)
		return
	}
	defer file.Close()

	// Create a scanner to read the file
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Bytes()
		fmt.Printf("Length: %d\n", len(line))
		fmt.Println(line)

		// find the location of each symbol in the line and store it in a map
		// a symbol can appear multiple times in a line and is not a "." and is not an int
		symbolMap := make(map[byte][]int)
		for i := 0; i < len(line); i++ {
			if line[i] != '.' && line[i] < '0' || line[i] > '9' {
				symbolMap[line[i]] = append(symbolMap[line[i]], i)
			}
		}

		fmt.Println(symbolMap)
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error: %v", err)
	}
}
