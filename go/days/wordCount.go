package main

import (
	"fmt"
	"strings"
)

func main() {

	wordCount := make(map[string]int)

	text := "We are going to go to the stadium. You are, too. We will see you there."

	words := strings.Split(text, " ")

	for _, word := range words {
		wordCount[word]++
	}

	fmt.Println(wordCount)
}
