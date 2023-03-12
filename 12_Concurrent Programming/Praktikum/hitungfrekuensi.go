package main

import (
	"fmt"
	"sync"
)

func main() {
	text := "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua"

	// Inisialisasi map untuk menyimpan frekuensi huruf
	freq := make(map[rune]int)

	// Inisialisasi WaitGroup untuk menunggu semua goroutine selesai
	var wg sync.WaitGroup

	// Membagi teks menjadi beberapa bagian
	parts := 4
	partSize := len(text) / parts

	// Loop untuk membuat dan menjalankan goroutine
	for i := 0; i < parts; i++ {
		// Menambahkan WaitGroup untuk setiap goroutine yang dibuat
		wg.Add(1)

		// Memecah teks menjadi bagian-bagian
		start := i * partSize
		end := (i + 1) * partSize
		if i == parts-1 {
			end = len(text)
		}
		partText := text[start:end]

		// Goroutine untuk menghitung frekuensi huruf pada bagian teks
		go func() {
			defer wg.Done()
			localFreq := make(map[rune]int)
			for _, r := range partText {
				if _, ok := localFreq[r]; !ok {
					localFreq[r] = 1
				} else {
					localFreq[r]++
				}
			}

			// Lock untuk mengakses map global
			// Menambahkan frekuensi huruf dari bagian teks ke map global
			for r, f := range localFreq {
				freqLock.Lock()
				freq[r] += f
				freqLock.Unlock()
			}
		}()
	}

	// Menunggu semua goroutine selesai
	wg.Wait()

	// Menampilkan frekuensi huruf yang telah dihitung
	for r, f := range freq {
		fmt.Printf("%c: %d\n", r, f)
	}
}

// Inisialisasi mutex lock untuk mengakses map frekuensi huruf global
var freqLock sync.Mutex
