package hangman

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Word struct { // Struct of the word
	The_word         string
	Word_runes       []rune
	N                int
	Attempts         int
	HangmanPositions [10]string
	Letter_used      []rune
	Alpha_letter     string
	Hard             bool
	Voyelle_used     []rune
}

// Infinite loop to play the game until the player wins or loses
func (w *Word) Play() {
	w.AfficheRune(w.Word_runes)
	fmt.Print("\n")
	for {
		if w.Attempts <= 0 { // Checks if the player has lost
			Print_lose()
			fmt.Println("The word was", w.The_word)
			return
		} else if w.Check_win() { // Checks if the player has won
			Print_win()
			fmt.Println("The word was", w.The_word)
			return
		} else {
			fmt.Println("Enter a letter")
			var letter rune
			input := bufio.NewScanner(os.Stdin) // Scan the input
			input.Scan()
			str := Convert_string_with_accent_to_string(input.Text()) // Accent management
			fmt.Println("Choose : " + strings.ToUpper(str))
			fmt.Print("\n")
			if len(str) == 1 {
				letter = rune(strings.ToLower(str)[0])
				AlphaSort(w.Letter_used)

				if w.Hard { // Hard mode active
					fmt.Println("hard mode")
					if w.Is_used(letter) {
						w.Attempts--
						fmt.Println("You have already used this letter, you have ", w.Attempts, " attempts remaining")
					} else {
						w.Letter_used = append(w.Letter_used, letter)
						w.Check_voyelle()
						w.Check_letter(letter)
					}
					// normal mode
				} else if w.Is_used(letter) { // Check if the letter has already been said
					fmt.Println(strings.ToUpper(input.Text()) + " is already used like : " + strings.ToUpper(string(w.Letter_used)))
				} else {
					w.Letter_used = append(w.Letter_used, letter)
					w.Check_letter(letter)
				}
			} else if len(str) > 1 { // Save the game
				if str == "STOP" {
					Save(*w)
					fmt.Println("Game saved in save.txt")
					return
				}
				if w.Check_word(strings.ToLower(str)) { // Check if the word is correct
					Print_win()
					fmt.Println("The word was", w.The_word)
					return
				} else {
					w.Attempts -= 2
					fmt.Println("Not present in the word, " + strconv.Itoa(w.Attempts) + " attemps remaining")
					if w.Attempts < 0 {
					} else {
						Affiche(9 - w.Attempts) // Print the hangman
					}
				}
			}
			w.AfficheRune(w.Word_runes) // Print the word
			fmt.Print("\n")
		}
	}
}
