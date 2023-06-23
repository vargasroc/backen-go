package main

import (
	
	"fmt"
)

/*
Registro de estudiantes
Una universidad necesita registrar a los estudiantes y generar 
una funcionalidad para imprimir el detalle de los datos de 
cada uno de ellos, de la siguiente manera:
Nombre: [Nombre del alumno]
Apellido: [Apellido del alumno]
DNI: [DNI del alumno]
Fecha: [Fecha ingreso alumno]
Los valores que están en corchetes deben ser reemplazados por 
los datos brindados por los alumnos. Para ello es necesario 
generar una estructura Alumno con las variables Nombre, 
Apellido, DNI, Fecha y que tenga un método detalle
*/

type Alumno struct {
Nombre string `Nombre:[Nombre del alumno]`
Apellido string `Apellido[Apellido del alumno]`
DNI string `[DNI del alumno]`
Fecha string`[Fecha ingreso alumno]`
}

func main() {
	a := Alumno{
		Nombre: "Gabriel",
		Apellido: "Gomez",
DNI: "123456789",
Fecha: "15/16/18",
	}

	a.Descripcion()
	
}

func (a *Alumno) Descripcion() {
	fmt.Printf("%s %s con DNI %s inicio clases el %s", a.Nombre, a.Apellido, a.DNI, a.Fecha)
}