package main

import (
	"errors"
	"fmt"
)

/*
Calcular cantidad de alimento
Un refugio de animales necesita calcular cuánto alimento debe
comprar para las mascotas. Por el momento solo tienen
tarántulas, hámsteres, perros y gatos, pero se espera que
puedan darle refugio a muchos animales más.
Tienen los siguientes datos:
Perro: 10 kg de alimento.
Gato: 5 kg de alimento.
Hamster: 250 g de alimento.
Tarántula: 150 g de alimento.
Se solicita:
Implementar una función animal que reciba como parámetro un
valor de tipo texto con el animal especificado y que retorne
una función y un mensaje (en caso de que no exista el animal).
Una función para cada animal que calcule la cantidad de
alimento basándose en la cantidad del tipo de animal
especificado.
Ejemplo:
const (
   dog    = "dog"
   cat    = "cat"
)
...
animalDog, err := animal(dog)
animalCat, err := animal(cat)
...
var amount float64
amount += animalDog(10)
amount += animalCat(10)
*/

func main() {
	animalPerro, err := animal("perro")
	if err != nil {
		fmt.Println(err)
		return
	}
	animalGato, err := animal("gato")
	if err != nil {
		fmt.Println(err)
		return
	}
	animalHamster, err := animal("hamster")
	if err != nil {
		fmt.Println(err)
		return
	}
	animalJulian, err := animal("julian")
	if err != nil {
		fmt.Println(err)
		return
	}

	var amount float64
	amount += animalPerro(1.00)
	amount += animalGato(10.60)
	amount += animalHamster(100.00)
	amount += animalJulian(109.0)

	fmt.Println("Total de comida que hay que comprar:", amount)

}

func animal(tipoAnimal string) (func(float64) float64, error) {

	switch tipoAnimal {
	case "perro":
		return func(cantidad float64) float64 {
			return cantidad * 10
		}, nil
	case "gato":
		return func(cantidad float64) float64 {
			return cantidad * 5
		}, nil
	case "hamster":
		return func(cantidad float64) float64 {
			return cantidad * 0.25
		}, nil
	case "julian":
		return func(cantidad float64) float64 {
			return cantidad * 0.15
		}, nil
	default:
		return nil, errors.New("Ese animal no lo tenemos")

	}
}
