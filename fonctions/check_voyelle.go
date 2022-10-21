package hangman

import "fmt"

// Vowel management
func (w *Word) Check_voyelle() bool{
	if w.Letter_used[len(w.Letter_used)-1] == 'a' || w.Letter_used[len(w.Letter_used)-1] == 'e' || w.Letter_used[len(w.Letter_used)-1] == 'i' || w.Letter_used[len(w.Letter_used)-1] == 'o' || w.Letter_used[len(w.Letter_used)-1] == 'u' || w.Letter_used[len(w.Letter_used)-1] == 'y' {
		w.Voyelle_used = append(w.Voyelle_used, w.Letter_used[len(w.Letter_used)-1])
		if len(w.Voyelle_used) > 3 {
			w.Attempts--
			fmt.Println("You already used 3 vowels you have", w.Attempts, "attempts left")
			return true
		}
	}
	return false
}
