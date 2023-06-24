package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	var filename string
	fmt.Scanln(&filename)

	words, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	taboo := strings.ToUpper(string(words))

	var word string
	fmt.Scanln(&word)

	for strings.ToUpper(word) != "EXIT" {

		if strings.Contains(taboo, strings.ToUpper(word)) {
			fmt.Println(strings.Repeat("*", len(word)))
		} else {
			fmt.Println(word)
		}

		fmt.Scanln(&word)
	}
	fmt.Println("Bye!")
}
