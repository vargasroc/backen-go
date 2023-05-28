package main

import "fmt"

//import "fmt"

/*
Ejercicio 1 - Impuestos de salario
Una empresa de chocolates necesita calcular el impuesto de
sus empleados al momento de depositar el sueldo.
Para cumplir el objetivo es necesario crear una función que
devuelva el impuesto de un salario, teniendo en cuenta que si
la persona gana más de $50.000 se le descontará un 17 % del
sueldo y si gana más de $150.000 se le descontará, además,
un 10 % (27 % en total).

*/



func main(){
	sueldo:= float32(165300.00)
	resultado:=impuesto(sueldo)
	fmt.Println(resultado)
}

func impuesto(sueldo float32) float32 {
	if sueldo >= 50000 && sueldo < 150000 {
		return sueldo - sueldo * 0.17
} else if sueldo >= 150000 {
	return sueldo - sueldo * 0.27
} 
	return sueldo
}



	
