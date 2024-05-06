package digitfinder

import (
	"fmt"
	"testing"
)

func TestFindDigits(t *testing.T) {
	line1 := "1asdfone3fivefour"
	wantFirst := 1

	foundDigits, _ := FindDigitsInLine(line1)

	fmt.Println(foundDigits.first)

	if foundDigits.first != wantFirst {
		t.Fatalf(`findDigitsInLine(%v) = %v want %v`, line1, foundDigits.first, wantFirst)
	}

	wantLast := 4
	if foundDigits.last != wantLast {
		t.Fatalf(`findDigitsInLine(%v) = %v want %v`, line1, foundDigits.last, wantLast)
	}
}
