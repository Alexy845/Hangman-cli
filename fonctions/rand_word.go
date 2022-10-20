package hangman

import (
	"runtime"
)

// Mot aléatoire dans un slice de mots
func Rand_word(file string) string {
	lst_words := Open_file(file, "\n")
	word := lst_words[Rand_range(0, len(lst_words)-1)]
	switch runtime.GOOS {// Selon le système d'exploitation
	case "windows":
		return word[:len(word)-1]
	case "linux":
		return word
	default: // Si le système d'exploitation est autre
		return word
	}
}
