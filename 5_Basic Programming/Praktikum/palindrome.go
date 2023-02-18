package main

import (
	"fmt"
	"strings"
)

func main() {
	var input string

	fmt.Println("Apakah palindrome ?")
	fmt.Print("Masukkan kata: ")
	fmt.Scanln(&input)

	input = strings.ToLower(input)

	palindrome := true
	Leng := len(input)

	for i := 0; i < Leng/2; i++ {
		if input[i] != input[Leng-i-1] {
			palindrome = false
			break
		}
	}

	fmt.Printf("Captured : %s", input)

	if palindrome {
		fmt.Printf("\n%s adalah kata palindrome.", input)
	} else {
		fmt.Printf("\n%s bukan kata palindrome.", input)
	}
}
