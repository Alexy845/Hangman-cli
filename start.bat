#!/bin/sh -x
@echo off

start_mode_normal() {
    echo "Starting in normal mode"
    # Start the application
    go run main.go asset/words.txt
}

start_mode_hard (){
    echo "Starting in hard mode"
    # Start the application
    go run main.go --hard asset/words.txt
}

start_with_save (){
    echo "Starting with save"
    # Start the application
    go run main.go --startWith save.txt
}
echo "___________________________"
echo "Choose a mode:            |"
echo "1. Normal                 |"
echo "2. Hard                   |"
echo "3. Start with save        |"
echo "___________________________"
read -p "Enter your choice: " choice

case $choice in
    1) start_mode_normal ;;
    2) start_mode_hard ;;
    3) start_with_save ;;
    *) echo "Invalid choice" ;;
esac
