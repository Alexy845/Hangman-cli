package hangman

import (
	"runtime"
)

// Select a random word in a file
func Rand_word(file string) string {
	lst_words := Open_file(file, "\n")
	word := lst_words[Rand_range(0, len(lst_words)-1)]
	switch runtime.GOOS {
	case "windows":
		return word[:len(word)-1]
	case "linux":
		return word
	default:
		return word
	}
}
