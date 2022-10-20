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
	objects_save := w
	save, err := json.Marshal(objects_save)
    if err != nil {
        fmt.Println(err)
    } 
	err = ioutil.WriteFile(".\\asset\\save.txt", save, 0600)
	if err != nil {
        fmt.Println(err)
    } 
}
