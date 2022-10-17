package hangman

import (
	"math/rand"
	"time"
)

// Random number in a range
func Rand_range(min int, max int) int {
    rand.Seed(time.Now().UnixNano())
    return rand.Intn(max-min) + min
}
