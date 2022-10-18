package hangman

// Choisit n lettres aléatoires dans le slice de runes à révéler
func (w *Word) Rand_letters() {
	var numbers int
	for i := 0; i < w.N; i++ { // Choisi n lettres aléatoires
		numbers = Rand_range(0, len(w.The_word)-1) // Choisi un nombre aléatoire
		if w.Word_runes[numbers] == 0 {            // Si la lettre n'a pas déjà été choisie
			w.Word_runes = Replace_rune(w.Word_runes, rune(w.The_word[numbers]), numbers) // Si la lettre n'est pas déjà dans le slice, on l'ajoute
		} else {
			i-- // Si la lettre est déjà dans le slice, on recommence
		}
	}
}
