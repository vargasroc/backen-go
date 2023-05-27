package main

import "fmt"

func main(){
	palabra:="fisioterapia"
	var cant int = len(palabra)

	fmt.Printf("La palabra %s tiene %d letras\n", palabra, cant)

}