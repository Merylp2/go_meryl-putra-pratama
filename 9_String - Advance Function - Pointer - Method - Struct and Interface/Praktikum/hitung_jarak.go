package main

import "fmt"

type Car struct {
	typeCar string
	fuelIn  int
}

func (car Car) hitungJarak() float64 {
	return float64(car.fuelIn) / (1.5)
}

func main() {
	car := Car{
		typeCar: "Avanza",
		fuelIn:  40,
	}
	fmt.Println("Jarak tempuh mobil", car.typeCar, "adalah", car.hitungJarak(), "mill")
}
