package main

// Deklarasi struct kendaraan
type Kendaraan struct {
	TotalRoda      int
	KecepatanMobil int
}

// Deklarasi struct mobil dari struct kendaraan
type Mobil struct {
	Kendaraan
}

// Deklarasi method tambahKecepatan dengan parameter kecepatanbaru bertipe int
func (mobil *Mobil) tambahKecepatan(kecepatanbaru int) {
	// Mengubah nilai kecepatan mobil berdasarkan kecepatan mobil sebelumnya ditambah kecepatan baru
	mobil.KecepatanMobil = mobil.KecepatanMobil + kecepatanbaru
}

// Membuat method berjalan untuk menambah kecepatan mobil
func (mobil *Mobil) berjalan() {
	mobil.tambahKecepatan(10)
}

func main() {
	// Deklarasi variabel MobilCepat dengan nilai Mobil
	MobilCepat := Mobil{Kendaraan{4, 60}}
	// Memanggil method berjalan agar kecepatan mobil bertambah
	MobilCepat.berjalan()
	MobilCepat.berjalan()
	MobilCepat.berjalan()

	// Deklarasi variabel MobilLamban dengan nilai Mobil
	MobilLamban := Mobil{Kendaraan{4, 40}}
	MobilLamban.berjalan()
}
