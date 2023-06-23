package main

import "fmt"

type Producto interface {
	Precio() float64
}

type Pequenio struct {
	costo float64
}

func (p Pequenio) Precio() float64 {
	return p.costo
}

type Mediano struct {
	costo float64
}

func (m Mediano) Precio() float64 {
	return m.costo + m.costo*0.03
}

type Grande struct {
	costo float64
}

func (g Grande) Precio() float64 {
	return g.costo + g.costo*0.06 + 2500
}

func ProductoFactory(tipo string, costo float64) Producto {
	switch tipo {
	case "PEQUENIO":
		return Pequenio{costo: costo}
	case "MEDIANO":
		return Mediano{costo: costo}
	case "GRANDE":
		return Grande{costo: costo}
	default:
		return nil
	}
}

func main() {
	productos := []Producto{
		ProductoFactory("PEQUENIO", 25.0),
		ProductoFactory("MEDIANO", 25.0),
		ProductoFactory("GRANDE", 25.0),
	}

	for _, producto := range productos {
		fmt.Println("Precio total:", producto.Precio())
	}
}
