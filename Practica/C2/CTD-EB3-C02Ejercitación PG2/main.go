package main

import "fmt"

/* 
Ejercicio 2 - Calcular promedio
Un colegio necesita calcular el promedio (por estudiante) 
de sus calificaciones. Se solicita generar una función en 
la cual se le pueda pasar N cantidad de enteros y devuelva 
el promedio. No se pueden introducir notas negativas.

 */



func main(){
	calificaciones := []int {5,9,6,8,9,6,8,5}
	totalPromedio:= promedio(calificaciones...)
	fmt.Println(totalPromedio)

}


func promedio(calificaciones...int)float64{
	totalCalificaciones := len(calificaciones)
	suma:=0
	for _, v := range calificaciones {
		
		suma += v
	}
	
	promedio:= float64(suma) / float64(totalCalificaciones)
	return promedio
		
	}
	
		



	
