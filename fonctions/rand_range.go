package hangman

import (
	"math/rand"
	"time"
)

// Nombre aléatoire dans un intervalle
func Rand_range(min int, max int) int {
	rand.Seed(time.Now().UnixNano()) // initialiser le générateur de nombres aléatoires
	return rand.Intn(max-min) + min  // retourner un nombre aléatoire dans l'intervalle
}
