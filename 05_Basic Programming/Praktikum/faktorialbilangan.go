package main

import "fmt"

func main() {
	angka := 26

	for i := 1; i <= angka; i++ {
		if angka%i == 0 {
			fmt.Println(i)
		}
	}
}