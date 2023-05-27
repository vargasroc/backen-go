package main

import "fmt"

/* 
Ejercicio 2 - Qué edad tiene...
Un empleado de una empresa quiere saber el nombre y edad de uno de sus empleados.
Según el siguiente map, debemos imprimir la edad de Benjamin.
var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}
Por otro lado, también es necesario:
■ Saber cuántos de sus empleados son mayores de 21 años.
■ Agregar un empleado nuevo a la lista, llamado Federico que tiene 25 años.
■ Eliminar a Pedro del map.
 */



func main(){
var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}
fmt.Println(employees["Benjamin"])
mayor21(employees)
delete(employees, "Pedro")
fmt.Println(employees)
}

func mayor21(employees map[string]int){
	mayores:= 0
	for _, edad := range employees {
		if edad > 21{
			mayores ++			
		}
		}
		fmt.Println("Cantidad de empleados mayores de 21 años:", mayores)
	}



	
