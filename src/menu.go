package main

import (
	"fmt"

	"github.com/kshksdrt/text-utils/features/describe"
)

func ShowMenu() {
	fmt.Print("\033[H\033[2J")
	fmt.Println("Enter your choice")
	fmt.Println("1. Get word count")
	var selection int
	fmt.Scanln(&selection)
	fmt.Print("\033[H\033[2J")
	switch selection {
	case 1:
		describe.WordCounter()
	}
}
