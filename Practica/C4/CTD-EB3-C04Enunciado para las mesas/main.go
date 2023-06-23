package main

import "fmt"

type Product struct {
	ID          int
	Name        string
	Price       float64
	Description string
	Category    string
}

var Products []Product

func main() {
	Products = []Product{
		{ID: 1, Name: "Producto 1", Price: 10.0, Description: "Descripción 1", Category: "Categoría 1"},
		{ID: 2, Name: "Producto 2", Price: 15.0, Description: "Descripción 2", Category: "Categoría 2"},
	}

	product := Product{ID: 3, Name: "Producto 3", Price: 20.0, Description: "Descripción 3", Category: "Categoría 3"}
	Save(product)

	GetAll()
	foundProduct := GetByID(2)
	if foundProduct != (Product{}) {
		fmt.Println("Producto encontrado:", foundProduct)
	} else {
		fmt.Println("Producto no encontrado.")
	}

	GetByID(3)
	

}

func Save(p Product) {
	Products = append(Products, p)
	fmt.Println("Producto guardado:", p)
}

func GetAll() {
	fmt.Println("Lista de productos:")
	for _, product := range Products {
		fmt.Println(product)
	}
}

func GetByID(ID int) Product {
	for _, product := range Products {
		if product.ID == ID {
			return product
		}
	}
	return Product{}
}