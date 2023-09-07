package common

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
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
		rand.Seed(time.Now().UnixNano())
		return rand.Intn(num)
	}
}

// GenerateRandomNumber returns a random number from
// a given max value as a int
func GenerateRandomNumber(max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max)
}

// GenerateRandomNumberRange returns a random number
// over a given minimum and maximum range.
func GenerateRandomNumberRange(min int, max int) int {
	return (rand.Intn(max-min) + min)
}
