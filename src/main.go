package main

import (
	"bufio"
	"fmt"
	"os"
	s "strings"
)

func main() {
	ShowMenu()
}

func ShowMenu() {
	fmt.Println("Enter your choice")
	fmt.Println("1. Get word count")
	var selection int
	fmt.Scanln(&selection)
	switch selection {
	case 1:
		fmt.Printf(string([]byte{0x1b, '[', '3', 'J'}))
		GetWordCount()
	}
}

func GetWordCount() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your text")
	text, _ := reader.ReadString('\n')
	text = s.Replace(text, "\n", "", -1)
	fmt.Println("Your text has", s.Count(s.Replace(text, "\n", "", -1), " ")+1, "words")
}
