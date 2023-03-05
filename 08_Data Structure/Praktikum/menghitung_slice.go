package main

import "fmt"

func Mapping(slice []string) map[string]int {
	// Membuat map kosong dan menampung di variabel result
	result := make(map[string]int)

	// Melakukan iterasi terhadap slice dan menambahkan nilai dari map result
	for _, value := range slice {
		result[value]++
	}
	return result
}

func main() {
	fmt.Println(Mapping([]string{"asd", "qwe", "asd"}))
}
