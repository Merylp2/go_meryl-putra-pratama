package main

import "fmt"

func ArrayMerge(arrayA, arrayB []string) []string {
	// Menggabungkan slice arrayA dan arrayB ke variabel merged
	merged := append(arrayA, arrayB...)
	// Membuat slice kosong dan menampung di variabel result
	result := make([]string, 0)
	// Membuat map kosong dan menampung di variabel set
	set := make(map[string]bool)

	for _, elem := range merged {
		// Jika elem belum ada di map set, maka akan menambahkan elem ke map set dan slice result
		if !set[elem] {
			// Menambahkan elem ke map set
			set[elem] = true
			// Menambahkan elem ke slice result
			result = append(result, elem)
		}
	}

	return result
}

func main() {
	fmt.Println(ArrayMerge([]string{"king", "devil jin", "akuma"}, []string{"eddie", "steve", "geese"}))
	fmt.Println(ArrayMerge([]string{"sergei", "jin"}, []string{"jin", "steve", "bryan"}))
	fmt.Println(ArrayMerge([]string{"alisa", "yoshimitsu"}, []string{"devil jin", "yoshimitsu", "alisa", "law"}))
	fmt.Println(ArrayMerge([]string{}, []string{"devil jin", "sergei"}))
	fmt.Println(ArrayMerge([]string{"hwoarang"}, []string{}))
	fmt.Println(ArrayMerge([]string{}, []string{}))

}
