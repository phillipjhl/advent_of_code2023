package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)

func main() {
	file_path := os.Args[1]
	file, err := os.Open(file_path)
    var	RuleMap = map[string]int{"red": 12, "green": 13, "blue": 14}
	var sumOfTotalGameIds int = 0

	if err != nil {
		fmt.Println("Error opening file")
		return
	}

	defer file.Close()

	// process each line of the file by reading each bufio
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		line = line[5:]
		var gameID string
		var impossibleFlag int = 0

		// loop over line to find the game id 
		for i := 0; i <= len(line)-1; i++ {
			// strip "Game {id}:" from each line, and store the id for the line
			character := line[i]
			if character == ':' {
				gameID = line[:i]
				line = line[i+2:]
				break
			}
		}

		listOfMatches := strings.SplitAfter(line, ";")
		// fmt.Println(listOfMatches)
		// loop over each match in the line
		for _, match := range listOfMatches {
			// fmt.Println(match)
			// strip the ";" from the match if it exists
			trimmedMatch := strings.TrimSuffix(match, ";")
			trimmedMatch = strings.TrimSpace(trimmedMatch)
			scores := strings.Split(trimmedMatch, ",")

			// loop over each score in the match and check if the score
			// is greater than the value in the RuleMap
			// if one rule is broken, invalidate the current game
			for _, score := range scores {
				score = strings.TrimSpace(score)
				scoreTuple := strings.Split(score, " ")

				cubeColor := scoreTuple[1]
				scoreInt, err := strconv.Atoi(scoreTuple[0])
				if err != nil {
					fmt.Println("Error converting score to int")
					return
				}
				for key, limit := range RuleMap {
					if  cubeColor == key && scoreInt > limit {
						fmt.Printf("Num Invalid for Game ID: %v, %v for %v over limit: %v \n", gameID, scoreInt, key, limit)
						// mark the current game is impossible
					    impossibleFlag = 1
						break
					}
				}
			}
		}

		if impossibleFlag == 0 {
			gameIDInt, err := strconv.Atoi(gameID)
			if err != nil {
				fmt.Println("Error converting score to int")
			}
			sumOfTotalGameIds += gameIDInt
		}

		fmt.Printf("Game ID: %v\n", gameID)
	}

	fmt.Println("Total Sum of valid Game IDs: ", sumOfTotalGameIds)
}
