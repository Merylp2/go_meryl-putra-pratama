package main

import "fmt"

func main() {
	angka := 5

	for i := 1; i <= angka; i++ {
		for s := i; s < angka; s++ {
			fmt.Print(" ")
		}
		for j := 1; j <= i; j++ {
			fmt.Print(" *")
		}
		fmt.Println()
	}
}
