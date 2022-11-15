package hangman

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

func (ww Word) Start() {

}
