package hangman

// Replace a rune with another
func Replace_rune(s []rune, r rune, i int) []rune {
	s[i] = r
	return s
}
