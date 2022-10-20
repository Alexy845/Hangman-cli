package hangman

import (
	"fmt"
	"io/ioutil"
)

// Verifie si y a une sauvegarde
func Is_not_empty () bool {
	save, err := ioutil.ReadFile(".\\asset\\save.txt")
	if err != nil {
		fmt.Println("Peut pas lire")
		panic(err)
	}
	if save != nil{
		return true
	}
	return false
}
	//var elements_save = []byte(`{"The_word":"","Word_runes":[],"N":3,"Attempts":10,"HangmanPositions":[],"Letter_used":[]}`)
	//https://stackoverflow.com/questions/24770403/write-struct-to-json-file-using-struct-fields-not-json-keys
	//err1 := json.Unmarshal(elements_save, &word)