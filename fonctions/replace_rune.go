package hangman

// Replace a rune in a slice of runes
func Replace_rune(s []rune, r rune, i int) []rune {
	s[i] = r
	return s
}