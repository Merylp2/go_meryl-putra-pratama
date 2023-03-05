package main

import (
	"fmt"
)

func primeX(number int) int {
	count := 0
	prime := 2

	for count < number {
		isPrime := true

		for i := 2; i < prime; i++ {
			if prime%i == 0 {
				isPrime = false
				break
			}
		}

		if isPrime {
			count++
		}

		prime++
	}

	return prime - 1
}

func main() {
	fmt.Println(primeX(1))  // 2
	fmt.Println(primeX(5))  // 11
	fmt.Println(primeX(8))  // 19
	fmt.Println(primeX(9))  // 23
	fmt.Println(primeX(10))	//29
}
