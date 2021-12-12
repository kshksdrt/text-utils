package describe

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func WordCounter() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your text")
	text, _ := reader.ReadString('\n')
	fmt.Println("Your text has", getWordCount(text), "words")
}

func getWordCount(text string) int {
	return strings.Count(strings.Replace(text, "\n", "", -1), " ") + 1
}
