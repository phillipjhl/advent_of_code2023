package digitfinder

import (
	"fmt"
	"testing"
)

func TestFindDigits(t *testing.T) {
	line1 := "1asdfone3fivefour"
	line2 := "72mmjrfjvlzone3threethreesix"
	given := [2]string{line1, line2}
	expected := 1
	expectedLast := [2]int{4, 6}

	for i := range given {
		foundDigits, _ := FindDigitsInLine(given[i])

		fmt.Printf("Found Digits %v, of given: %v \n", foundDigits, given[i])

		if foundDigits.First != expected {
			t.Fatalf(`findDigitsInLine(%v) = %v want %v`, given[i], foundDigits.First, expected)
		}

		if foundDigits.Last != expectedLast[i] {
			t.Fatalf(`findDigitsInLine(%v) = %v want %v`, given[i], foundDigits.Last, expectedLast[i])
		}
	}
}
