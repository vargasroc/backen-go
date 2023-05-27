package main

import "fmt"

/* Ejercicio 1 - Descuento
Una tienda de ropa quiere ofrecer a sus clientes un descuento
sobre sus productos. Para ello necesitan una aplicación que 
les permita calcular el descuento basándose en dos variables: 
su precio y el descuento en porcentaje. La tienda espera 
obtener como resultado el valor con el descuento aplicado y 
luego imprimirlo en consola.
 */



func main(){
precio:= 150.36
porcentaje:= 0.36
precioTotal:= calcularPrecioTotal(precio, porcentaje)
fmt.Println(precioTotal)
}

func calcularPrecioTotal(precio float64, porcentaje float64) float64 {
	precioTotal := precio -  (precio* porcentaje)
	return precioTotal
	
}