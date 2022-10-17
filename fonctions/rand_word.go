package hangman

// Mot aléatoire dans un slice de mots
func Rand_word(file string) string {
	lst_words := Open_file(file, "\n")                 // ouvrir le fichier et récupérer les mots
	word := lst_words[Rand_range(0, len(lst_words)-1)] // récupérer un mot aléatoire
	return word                                        // retourner le mot
}
