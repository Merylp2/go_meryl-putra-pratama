package main

import "fmt"

func PairSum(arr []int, target int) []int {
	// Membuat variabel left dan right untuk menampung index pertama dan terakhir dari slice
	left, right := 0, len(arr)-1

	// Melakukan perulangan selama left lebih kecil dari right
	for left < right {
		// Membuat variabel sum untuk menampung hasil penjumlahan dari nilai arr[left] dan arr[right]
		sum := arr[left] + arr[right]
		// Jika nilai sum sama dengan target, maka akan mengembalikan nilai left dan right
		if sum == target {
			// Mengembalikan nilai left dan right
			return []int{left, right}
			// Jika nilai sum lebih kecil dari target, maka akan menambahkan nilai left dengan 1
		} else if sum < target {
			// Menambahkan nilai left dengan 1
			left++
			// Jika nilai sum lebih besar dari target, maka akan mengurangi nilai right dengan 1
		} else {
			// Mengurangi nilai right dengan 1
			right--
		}
	}
	// Mengembalikan nilai slice kosong
	return []int{}
}

func main() {
	fmt.Println(PairSum([]int{1, 2, 3, 4, 6}, 6))
	fmt.Println(PairSum([]int{2, 5, 9, 11}, 11))
	fmt.Println(PairSum([]int{1, 3, 5, 7}, 12))
	fmt.Println(PairSum([]int{1, 4, 6, 8}, 10))
	fmt.Println(PairSum([]int{1, 5, 6, 7}, 6))
}
