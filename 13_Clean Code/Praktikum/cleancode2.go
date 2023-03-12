package main

type Kendaraan struct {
	TotalRoda      int
	KecepatanMobil int
}

type Mobil struct {
	Kendaraan
}

func (mobil *Mobil) tambahKecepatan(kecepatanbaru int) {
	mobil.KecepatanMobil = mobil.KecepatanMobil + kecepatanbaru
}

func (mobil *Mobil) berjalan() {
	mobil.tambahKecepatan(10)
}

func main() {
	MobilCepat := Mobil{Kendaraan{4, 60}}
	MobilCepat.berjalan()
	MobilCepat.berjalan()
	MobilCepat.berjalan()

	MobilLamban := Mobil{Kendaraan{4, 40}}
	MobilLamban.berjalan()
}