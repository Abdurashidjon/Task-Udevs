package main

import (
	"fmt"
	"strings"
)

func main() {
	a := Palindrom("AZiza")
	fmt.Println(a)
}

func Palindrom(word string) bool {
	convertUpper := strings.ToUpper(word)
	for i := 0; i < len(convertUpper); i++ {
		j := len(convertUpper) - 1 - i
		if convertUpper[i] != convertUpper[j] {
			return false
		}
	}
	return true
}
