package main

import "fmt"

/* 
Ejercicio 2 - Préstamo
Un banco quiere otorgar préstamos a sus clientes, pero no 
todos pueden acceder a los mismos. El banco tiene ciertas 
reglas para saber a qué cliente se le puede otorgar: solo le
otorga préstamos a clientes cuya edad sea mayor a 22 años, 
se encuentren empleados y tengan más de un año de antigüedad 
en su trabajo. Dentro de los préstamos que otorga, no
les cobrará interés a los que su sueldo sea mejor a $100.000.
Es necesario realizar una aplicación que tenga estas variables
y que imprima un mensaje de acuerdo a cada caso.
Pista: Tu código tiene que poder imprimir al menos tres 
mensajes diferentes.
 */



func main(){
edad:= 19
empleado:= true
antiguedad:= 5
sueldo:= 90000
accedeAprestamo(edad,empleado ,antiguedad,sueldo)

}

func accedeAprestamo(edad int, empleado bool, antiguedad int, sueldo int) {
	if edad >= 22 && empleado == true && antiguedad >= 1 &&  sueldo >=100000{
		fmt.Println("accedes a prestamo pero CON intereses")
	} else if edad >= 22 && empleado == true && antiguedad >= 1 && sueldo <=100000 {
		fmt.Println("accedes a prestamo pero SIN intereses")
	} else{
	 fmt.Println("sali de aca pibe")
	}
}

	
