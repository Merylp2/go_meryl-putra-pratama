package main

import (
	"fmt"
	"time"
)

func kelipatan(x int) {
	for i := 1; ; i++ {
		if i%x == 0 {
			fmt.Println(i)
			time.Sleep(3 * time.Second)
		}
	}
}

func main() {
	go kelipatan(3)

	// Akan berhenti jika sudah mencapai 30 detik
	time.Sleep(6 * time.Second)
}
