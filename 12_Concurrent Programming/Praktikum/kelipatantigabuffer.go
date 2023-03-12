package main

import (
	"fmt"
	"time"
)

func printMultiples(x chan int) {
	for i := 1; ; i++ {
		if i%3 == 0 {
			x <- i
			time.Sleep(3 * time.Second)
		}
	}
}

func main() {
	var buffer int = 10
	x := make(chan int, buffer) // buffer channel dengan kapasitas 10

	go printMultiples(x)

	for i := 0; i < buffer; i++ {
		fmt.Println(<-x)
	}
}
