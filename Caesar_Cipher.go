package main

import (
	"fmt"
	"strings"
	"unicode"
)

func CaesarCipherList(m []string, shift int, chann chan string) {
	s := make([]string, len(m))
	for i, word := range m {
		M := strings.ToUpper(word)
		for _, char := range M {
			if unicode.IsLetter(char) {
				s[i] += string(rune(((int(char) + shift - int('A')) % 26) + int('A')))
			}

		}
		chann <- s[i]

	}
}
func main() {
	messages := []string{"Csi2520", "Csi2120", "3 Paradigms",
		"Go is 1st", "Prolog is 2nd", "Scheme is 3rd",
		"uottawa.ca", "csi/elg/ceg/seg", "800 King Edward"}
	chann := make(chan string, len(messages))
	//	go CaesarCipherList(messages[:], 2, chann)
	go CaesarCipherList(messages[:int(len(messages)/3)], 2, chann)
	go CaesarCipherList(messages[int(len(messages)/3):(int(len(messages)/3)*2)], 2, chann)
	go CaesarCipherList(messages[(int(len(messages)/3)*2):], 2, chann)

	for range messages {
		fmt.Println(<-chann)
	}
}
