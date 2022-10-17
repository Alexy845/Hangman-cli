package hangman 

// Random words in a slice
func Rand_word(file string) string {
	lst_words := Open_file(file, "\n")
	word := lst_words[Rand_range(0, len(lst_words)-1)]
	return word
}