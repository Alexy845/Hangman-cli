package hangman

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Word struct {
	The_word         string
	Word_runes       []rune
	N                int
	Attempts         int
	HangmanPositions [10]string
}

// Choose n random letters in a slice of runes
func (w *Word) Rand_letters() {
	var numbers int
	for i := 0; i < w.N; i++ {
		numbers = Rand_range(0, len(w.The_word)-1)
		if w.Word_runes[numbers] == 0 {
			w.Word_runes = Replace_rune(w.Word_runes, rune(w.The_word[numbers]), numbers)
		} else {
			i--
		}
	}
}

// infinite loop to play the game until the player wins or loses
func (w *Word) Play() {
	fmt.Println("Good Luck, you have 10 attempts.")
	fmt.Println(strings.ToUpper(AfficheRune(w.Word_runes)))
	fmt.Print("\n")
	for {
		if w.Attempts == 0 {
			fmt.Println("You lost !")
			fmt.Println("The word was", w.The_word)
			break
		} else if w.Check_win() {
			fmt.Println("Congrat !")
			fmt.Println("The word was", w.The_word)
			break
		} else {
			fmt.Println("Enter a letter")
			var letter rune
			input := bufio.NewScanner(os.Stdin) // créer un scanner sur l'entrée
			input.Scan()                        // lance le scan
			for _, i := range input.Text() {
				letter = i
			}
			fmt.Println("Choose : " + strings.ToUpper(input.Text()))
			fmt.Print("\n")
			w.Check_letter(letter)
			fmt.Println(strings.ToUpper(AfficheRune(w.Word_runes)))
			fmt.Print("\n")
		}
	}
}
