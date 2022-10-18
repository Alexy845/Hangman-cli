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
	Letter_used      []rune
}

// infinite loop to play the game until the player wins or loses
func (w *Word) Play() {
	fmt.Println("Good Luck, you have 10 attempts.")
	fmt.Println(strings.ToUpper(AfficheRune(w.Word_runes)))
	fmt.Print("\n")
	for {
		if w.Attempts < 0 { // Si le joueur n'a plus d'essais il perd
			Print_lose()
			fmt.Println("The word was", w.The_word)
			return // On sort de la boucle
		} else if w.Check_win() { // Si le joueur a trouvé le mot il gagne
			Print_win()
			fmt.Println("The word was", w.The_word)
			return // On sort de la boucle
		} else {
			fmt.Println("Enter a letter")
			var letter rune                     // Création d'une variable rune
			input := bufio.NewScanner(os.Stdin) // Créer un scanner sur l'entrée
			input.Scan()                        // Lance le scan
			str := Convert_string_with_accent_to_string(input.Text())
			fmt.Println("Choose : " + strings.ToUpper(str))
			fmt.Print("\n")
			if len(str) == 1 { // Si la l'entrée est une lettre
				letter = rune(strings.ToLower(str)[0])        // Converti la lettre en rune
				w.Check_letter(letter)                        // Verifie si la lettre est dans le mot
				w.Letter_used = append(w.Letter_used, letter) // On ajoute la lettre dans le slice des lettres utilisées
			} else if len(str) > 1 { // Si la longueur de l'entrée est supérieur à 1
				if str == "STOP" || str == "stop" {
					Save(*w)
					return
				}
				if w.Check_word(strings.ToLower(str)) { // Verifie si le mot est le bon
					Print_win()
					fmt.Println("The word was", w.The_word)
					return // Si le mot est bon, on sort de la boucle
				} else {
					w.Attempts -= 2 // sinon perd 2 vies
					fmt.Println("Not present in the word, " + strconv.Itoa(w.Attempts) + " attemps remaining")
					if w.Attempts < 0 {
					} else {
						Affiche(9 - w.Attempts) // On affiche le pendu
					}
				}
			}
			fmt.Println(strings.ToUpper(AfficheRune(w.Word_runes))) // Affiche le mot avec les lettres trouvées
			fmt.Print("\n")
		}
	}
}
