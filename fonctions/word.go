package hangman

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Word struct { // Structure du mot à trouver
	The_word         string
	Word_runes       []rune
	N                int
	Attempts         int
	HangmanPositions [10]string
	OS               string
}

// Choose n random letters in a slice of runes
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

// infinite loop to play the game until the player wins or loses
func (w *Word) Play() {
	fmt.Println("Good Luck, you have 10 attempts.")
	fmt.Println(w.The_word)
	fmt.Println(strings.ToUpper(AfficheRune(w.Word_runes)))
	fmt.Print("\n")
	for {
		if w.Attempts == 0 { // Si le joueur n'a plus d'essais il perd
			fmt.Println("You lost !")
			fmt.Println("The word was", w.The_word)
			return // On sort de la boucle
		} else if w.Check_win() { // Si le joueur a trouvé le mot il gagne
			fmt.Println("Congrat !")
			fmt.Println("The word was", w.The_word)
			return // On sort de la boucle
		} else {
			fmt.Println("Enter a letter")
			var letter rune                     // Création d'une variable rune
			input := bufio.NewScanner(os.Stdin) // Créer un scanner sur l'entrée
			input.Scan()                        // Lance le scan
			fmt.Println("Choose : " + strings.ToUpper(input.Text()))
			fmt.Print("\n")
			if len(input.Text()) == 1 { // Si la l'entrée est une lettre
				letter = rune(strings.ToLower(input.Text())[0]) // Converti la lettre en rune
				w.Check_letter(letter)         // Verifie si la lettre est dans le mot
			} else if len(input.Text()) > 1 { // Si la longueur de l'entrée est supérieur à 1
				if w.Check_word(strings.ToLower(input.Text())) { // Verifie si le mot est le bon
					fmt.Println("Congrat !")
					fmt.Println("The word was", w.The_word)
					return // Si le mot est bon, on sort de la boucle
				} else { // Sinon on perd une vie
					w.Attempts -= 2 // On perd 2 vies
					fmt.Println("Not present in the word, " + strconv.Itoa(w.Attempts) + " attemps remaining")
					Affiche(9 - w.Attempts) // On affiche le pendu
				}
			}
			fmt.Println(strings.ToUpper(AfficheRune(w.Word_runes))) // Affiche le mot avec les lettres trouvées
			fmt.Print("\n")
		}
	}
}
