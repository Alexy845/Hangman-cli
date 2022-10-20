package hangman

import (
	"math/rand"
	"time"
)

// Nombre aléatoire dans un intervalle
func Rand_range(min int, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min) + min
}
