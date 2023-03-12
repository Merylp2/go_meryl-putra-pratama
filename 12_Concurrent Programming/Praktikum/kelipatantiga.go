package main

import (
	"fmt"
	"time"
)

func kelipatantiga(x chan int) {
	for i := 1; ; i++ {
		if i%3 == 0 {
			x <- i
			time.Sleep(3 * time.Second)
		}
	}
}

func main() {
	x := make(chan int)
	go kelipatantiga(x)

	for i := 0; i < 10; i++ {
		fmt.Println(<-x)
	}
}
