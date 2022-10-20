package hangman

// remplacer une rune dans un tableau de rune
func Replace_rune(s []rune, r rune, i int) []rune {
	s[i] = r
	return s
}
