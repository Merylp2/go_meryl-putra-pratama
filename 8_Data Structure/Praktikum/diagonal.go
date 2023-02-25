package main

import (
	"fmt"
	"math"
)

func diagonalDifference(arr [][]int32) int32 {
	// Menghitung jumlah diagonal kiri ke kanan
	diagonal1 := int32(0)
	for i := 0; i < len(arr); i++ {
		diagonal1 += arr[i][i]
	}

	// Menghitung jumlah diagonal kanan ke kiri
	diagonal2 := int32(0)
	for i := 0; i < len(arr); i++ {
		diagonal2 += arr[i][len(arr)-i-1]
	}

	// Menghitung selisih absolut antara jumlah diagonal
	diff := int32(math.Abs(float64(diagonal1 - diagonal2)))

	// Mengembalikan hasil
	return diff
}

func main() {
	// Mendefinisikan array 2 dimensi
	arr := [][]int32{
		{1, 2, 3},
		{4, 5, 6},
		{9, 8, 9}}

	// Memanggil fungsi diagonalDifference untuk menghitung selisih absolut antara jumlah diagonal
	diff := diagonalDifference(arr)

	// Menampilkan hasil
	fmt.Println(diff)
}
