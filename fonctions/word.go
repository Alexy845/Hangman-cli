package hangman

import (
	"fmt"
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

func (ww *Word) Start() {

	fmt.Println("Good Luck, you have 10 attempts.")

	ww.N = len(ww.The_word)/2 - 1
	ww.Word_runes = make([]rune, len(ww.The_word))
	ww.Rand_letters()
	println(ww.The_word)

	/*default: // If there is 3 args or more
		fmt.Print("Too many arguments !")
	}*/
}
