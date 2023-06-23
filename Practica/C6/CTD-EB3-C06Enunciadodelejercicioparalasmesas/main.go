package main

import "fmt"

type PaymentMethod interface {
	Pay(purchase *Purchase)
}

type Purchase struct {
	precio         float64
	PaymentMethods []PaymentMethod
}

type Credito struct{}

func (c Credito) Pay(purchase *Purchase) {
	purchase.precio = purchase.precio * 1.1
}

type Bancaria struct{}

func (b Bancaria) Pay(purchase *Purchase) {
	// No se aplica ningún cambio en el precio
}

type Efectivo struct{}

func (e Efectivo) Pay(purchase *Purchase) {
	purchase.precio = purchase.precio * 0.95
}

func main() {
	var precio float64
	fmt.Print("Ingrese el precio de la compra: ")
	fmt.Scanln(&precio)

	purchase := Purchase{
		precio: precio,
		PaymentMethods: []PaymentMethod{
			Credito{},
			Bancaria{},
			Efectivo{},
		},
	}

	var metodoPago string
	fmt.Print("Seleccione el método de pago (CREDITO, BANCARIA, EFECTIVO): ")
	fmt.Scanln(&metodoPago)

	for _, method := range purchase.PaymentMethods {
		switch metodoPago {
		case "CREDITO":
			if _, ok := method.(Credito); ok {
				method.Pay(&purchase)
			}
		case "BANCARIA":
			if _, ok := method.(Bancaria); ok {
				method.Pay(&purchase)
			}
		case "EFECTIVO":
			if _, ok := method.(Efectivo); ok {
				method.Pay(&purchase)
			}
		}
	}

	fmt.Println("Precio de la compra:", purchase.precio)
}
