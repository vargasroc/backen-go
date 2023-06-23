package main
import "fmt"

/*Una empresa de redes sociales requiere implementar una estructura usuarios con funciones que vayan agregando información a la misma. 
Para optimizar y ahorrar memoria requieren que la estructura usuarios ocupe el mismo lugar en memoria para el main del programa y para 
las funciones. La estructura debe tener los campos: nombre, apellido, edad, correo y contraseña. Y deben implementarse las funciones:
cambiarNombre: permite cambiar el nombre y apellido.
cambiarEdad: permite cambiar la edad.
cambiarCorreo: permite cambiar el correo.
cambiarContraseña: permite cambiar la contraseña.
*/

type usuarios struct{
	nombre string
	apellido string
	edad int
	correo string
	contrasenia string
}

func (u *usuarios) cambiarNombre (nuevoNombre, nuevoApellido string){
	u.nombre = nuevoNombre
	u.apellido = nuevoApellido

}
func (u *usuarios) cambiarEdad(nuevaEdad int){
	u.edad = nuevaEdad
}

func (u *usuarios) cambiarCorreo(nuevoCorreo string){
	u.correo = nuevoCorreo

}

func (u *usuarios) cambiarContrasenia(nuevaContrasenia string){
	u.contrasenia = nuevaContrasenia
}


func main(){

	usuarioCopado := usuarios{
		nombre:      "Rocio",
		apellido:    "Vargas",
		edad:        35,
		correo:      "vargasroc@gmail.com",
		contrasenia: "1234abc",
	}

	fmt.Println(usuarioCopado)

	usuarioCopado.cambiarNombre("Roc", "Var")
	usuarioCopado.cambiarEdad(18)
	usuarioCopado.cambiarContrasenia("9876poi")
	usuarioCopado.cambiarCorreo("rocbuscatrabajo@gmail.com")

	fmt.Println(usuarioCopado)

	}

