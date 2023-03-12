package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Product struct {
	Title    string  `json:"title"`
	Price    float64 `json:"price"`
	Category string  `json:"category"`
}

func fetchProduct(ch chan []Product) {
	res, err := http.Get("https://fakestoreapi.com/products")
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	var products []Product
	err = json.NewDecoder(res.Body).Decode(&products)

	if err != nil {
		panic(err)
	}

	ch <- products
}

func main() {
	ch := make(chan []Product)
	go fetchProduct(ch)

	products := <-ch

	fmt.Println("Products Data")
	for _, product := range products {
		fmt.Println("==============")
		fmt.Println("Title: ", product.Title)
		fmt.Println("Price: ", product.Price)
		fmt.Println("Category: ", product.Category)
	}
}
