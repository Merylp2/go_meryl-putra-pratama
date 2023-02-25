package main

import (
	"fmt"
	"strconv"
)

func munculSekali(angka string) []int {
	// Membuat map kosong dengan variable counts
	counts := make(map[int]int)

	// Melakukan perulangan terhadap string angka
	for _, char := range angka {
		// Mengubah string ke integer
		num, err := strconv.Atoi(string(char))
		// Jika terjadi error, maka akan melanjutkan ke perulangan selanjutnya
		if err != nil {
			continue
		}
		// Menambahkan nilai dari map counts dengan key num
		counts[num]++
	}

	// Membuat variable result dengan tipe data slice integer
	var result []int
	// Melakukan iterasi terhadap map counts
	for num, count := range counts {
		// Jika nilai count sama dengan 1, maka akan menambahkan nilai num ke slice result
		if count == 1 {
			// Menambahkan nilai num ke slice result
			result = append(result, num)
		}
	}

	return result
}

func main() {
	fmt.Println(munculSekali("1231234"))
}
