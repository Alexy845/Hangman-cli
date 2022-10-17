package hangman

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Word struct { // Structure du mot à trouver
	The_word         string
	Word_runes       []rune
	N                int
	Attempts         int
	HangmanPositions [10]string
}

// Choisit n lettres aléatoirement et les affiches à la place des _ dans le slice de runes
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

// Boucle de jeu principale infinie jusqu'à ce que le joueur gagne ou perde
func (w *Word) Play() {
	fmt.Println("Good Luck, you have 10 attempts.")
	fmt.Println(strings.ToUpper(AfficheRune(w.Word_runes)))
	fmt.Print("\n")
	for {
		if w.Attempts == 0 { // Si le joueur n'a plus d'essais il perd
			fmt.Println("You lost !")
			fmt.Println("The word was", w.The_word)
			break
		} else if w.Check_win() { // Si le joueur a trouvé le mot il gagne
			fmt.Println("Congrat !")
			fmt.Println("The word was", w.The_word)
			break
		} else {
			fmt.Println("Enter a letter")
			var letter rune
			input := bufio.NewScanner(os.Stdin)               // Créer un scanner sur l'entrée
			input.Scan()                                      // Lance le scan
			for _, i := range strings.ToLower(input.Text()) { // Si la lettre est en majuscule, on la passe en minuscule
				letter = i
			}
			fmt.Println("Choose : " + strings.ToUpper(input.Text()))
			fmt.Print("\n")
			w.Check_letter(letter)                                  // Verifie si la lettre est dans le mot
			fmt.Println(strings.ToUpper(AfficheRune(w.Word_runes))) // Affiche le mot avec les lettres trouvées
			fmt.Print("\n")
		}
	}
}
