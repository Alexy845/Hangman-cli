package hangman

import (
	"fmt"
)
// Vérifie si le mot est trouvé
func (w *Word) Check_word(input string) bool {
	fmt.Println(input)
	fmt.Println(len(input))
	fmt.Println(w.The_word)
	fmt.Println(len(w.The_word))
	return w.The_word == input // Retourne vrai si le mot est trouvé et faux sinon
}
