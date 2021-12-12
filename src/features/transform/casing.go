package transform

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func SentenceCaseConverter() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your text")
	text, _ := reader.ReadString('\n')
	fmt.Println(ToSentenceCase(text))
}

func ToSentenceCase(text string) string {
	return strings.ToUpper(text[0:1]) + strings.ToLower(text[1:])
}
