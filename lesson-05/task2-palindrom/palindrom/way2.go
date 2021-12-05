package main

import (
	"fmt"
	"strings"
)

func main() {
	a := Palindrom2("AziZa")
	fmt.Println(a)
}

func Palindrom2(word string) bool {
	convertWord := strings.ToUpper(word)
	newWord := ""
	for i := len(convertWord) - 1; i >= 0; i-- {
		newWord += string(convertWord[i])
	}
	for i := range word {
		if convertWord[i] != newWord[i] {
			return false
		}
	}
	return true
}
