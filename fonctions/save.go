package hangman

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func Save(w Word) {
	_, err := os.OpenFile(".\\asset\\save.txt", os.O_CREATE, 0600) 
	if err != nil {            
		panic(err) 
	}
	objects_save := w //[]string{w.The_word, w.Word_runes, w.N, w.Attempts, w.HangmanPositions, w.Letter_used}
	save, err := json.Marshal(objects_save)
     
    if err != nil {
        fmt.Println(err)
    } 
	err = ioutil.WriteFile(".\\asset\\save.txt", save, 0600)
}
