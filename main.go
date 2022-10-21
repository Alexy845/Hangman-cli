package main

import (
	"encoding/json"
	"fmt"
	hangman "hangman/fonctions"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

func main() {
	switch len(os.Args) {
	case 1: // If there isn't args
		fmt.Print("Please enter dictionary file")
		return
	case 2: // If there is one arg = normal game
		dico := os.Args[1]
		if _, err := os.Stat(os.Args[1]); err != nil {
			log.Fatal(err)
		}
		word := hangman.Word{The_word: hangman.Rand_word(dico), Word_runes: []rune{}, N: 0, Attempts: 10, Letter_used: []rune{}, Alpha_letter: "abcdefghijklmnopqrstuvwxyz", Hard: false}
		word.N = len(word.The_word)/2 - 1
		word.Word_runes = make([]rune, len(word.The_word))
		word.Rand_letters()
		fmt.Println("Good Luck, you have 10 attempts.")
		word.Play()
	case 3: // If there is 2 args = game from save.txt or --hard flags
		if os.Args[1] == "--hard" {
			dico := os.Args[2]
			if _, err := os.Stat(os.Args[2]); err != nil {
				log.Fatal(err)
			}
			word := hangman.Word{The_word: hangman.Rand_word(dico), Word_runes: []rune{}, N: 0, Attempts: 10, Letter_used: []rune{}, Alpha_letter: "abcdefghijklmnopqrstuvwxyz", Hard: true}
			word.N = len(word.The_word)/3 - 1
			word.Word_runes = make([]rune, len(word.The_word))
			word.Rand_letters()
			fmt.Println("Good Luck, you have 10 attempts.")
			word.Play()
		} else {
			fmt.Println("Please enter a valid argument")
		}
		if os.Args[1] == "--startWith" {
			if os.Args[2] == "save.txt" {
				if hangman.Is_not_empty() {
					var jsonS, err = ioutil.ReadFile(".\\asset\\" + os.Args[2])
					if err != nil {
						panic(err)
					}
					var object_save hangman.Word
					err = json.Unmarshal([]byte(jsonS), &object_save)
					if err != nil {
						panic(err)
					}
					fmt.Println("Welcome Back, you have : " + strconv.Itoa(object_save.Attempts) + " attemps remaining")
					object_save.Play()
				} else {
					fmt.Println("No save found ")
				}
			}
		}
	default: // If there is 3 args or more
		fmt.Print("Too many arguments !")
	}
}
