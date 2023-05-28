package main

import "fmt"

/* 
Ejercicio 1 - Calcular salario
Una empresa marinera necesita calcular el salario de sus 
empleados basándose en la cantidad de horas trabajadas 
por mes y la categoría.
Categoría C: su salario es de $1.000 por hora.
Categoría B: su salario es de $1.500 por hora, más un 20 % 
de su salario mensual.
Categoría A: su salario es de $3.000 por hora, más un 50 % 
de su salario mensual.
Se solicita generar una función que reciba por parámetro la 
cantidad de minutos trabajados por mes y la categoría, y que 
devuelva su salario.

 */


 

func main(){
	minutos:= 6000
	categoria:= "A" 
	salario(minutos, categoria)
	sueldoCompleto := salario(minutos, categoria)
	fmt.Println(sueldoCompleto)

}


func salario(minutos int, categoria string) float64{
	var horas = minutos /60 
	switch categoria {
	case "A":
		var sueldo = float64(horas * 3000)
		var sueldoCompleto = sueldo + sueldo * 0.5 
		return sueldoCompleto
	case "B": 
		var sueldo = float64(horas * 1500)
		var sueldoCompleto = sueldo + sueldo * 0.2 
		return sueldoCompleto
	case "C":
		var sueldo = float64(horas * 1000)
		return sueldo
	default:
return float64(minutos )
}
	
}