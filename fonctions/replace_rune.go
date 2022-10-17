package hangman

// Replace a rune in a slice of runes
func Replace_rune(s []rune, r rune, i int) []rune {
	s[i] = r // remplacer le caractère à l'index i par le caractère r
	return s // retourner le slice de runes
}
