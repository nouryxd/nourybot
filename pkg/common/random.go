package common

import (
	"fmt"
	"math/rand"
	"strconv"
)

// StrGenerateRandomNumber generates a random number from
// a given max value as a string
func StrGenerateRandomNumber(max string) int {
	num, err := strconv.Atoi(max)
	if num < 1 {
		return 0
	}

	if err != nil {
		fmt.Printf("Supplied value %v is not a number", num)
		return 0
	} else {
		return rand.Intn(num)
	}
}

// GenerateRandomNumber returns a random number from
// a given max value as a int
func GenerateRandomNumber(max int) int {
	return rand.Intn(max)
}

// GenerateRandomNumberRange returns a random number
// over a given minimum and maximum range.
func GenerateRandomNumberRange(min int, max int) int {
	return (rand.Intn(max-min) + min)
}
