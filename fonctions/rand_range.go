package hangman

import (
	"math/rand"
	"time"
)

// Generate a random int
func Rand_range(min int, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min) + min
}
