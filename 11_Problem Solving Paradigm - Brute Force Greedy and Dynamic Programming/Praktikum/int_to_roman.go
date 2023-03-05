package main

import "fmt"

func IntToRoman(num int) string {
	var roman string = ""
	var numbers = []int{1, 4, 5, 9, 10, 40, 50, 90, 100, 400, 500, 900, 1000}
	var romans = []string{"I", "IV", "V", "IX", "X", "XL", "L", "XC", "C", "CD", "D", "CM", "M"}
	var index = len(romans) - 1

	for num > 0 {
		for numbers[index] <= num {
			roman += romans[index]
			num -= numbers[index]
		}
		index -= 1
	}

	return roman
}

func main() {
	fmt.Println(IntToRoman(4))		// IV
	fmt.Println(IntToRoman(9))		// IX
	fmt.Println(IntToRoman(23))		// XXIII
	fmt.Println(IntToRoman(2021))	// MMXXI
	fmt.Println(IntToRoman(1646))	// MDCXLVI
}
