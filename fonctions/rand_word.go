package hangman

import (
	"runtime"
)

// Mot aléatoire dans un slice de mots
func Rand_word(file string) string {
	lst_words := Open_file(file, "\n")                 // ouvrir le fichier et récupérer les mots
	word := lst_words[Rand_range(0, len(lst_words)-1)] // récupérer un mot aléatoire
	switch runtime.GOOS {                              // Selon le système d'exploitation
	case "windows": // Si le système d'exploitation est windows
		return word[:len(word)-1] // On retourne word -1
	case "linux": // Si le système d'exploitation est linux
		return word // On retourne word
	default: // Si le système d'exploitation est autre
		return word // On retourne word
	}
}
