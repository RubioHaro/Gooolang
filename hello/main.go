package main

import "fmt"

func main() {

	var mensaje string = "Hola mundo"
	mensajeFacil := "Hola mundo con :="

	fmt.Println(mensaje)
	fmt.Println(mensajeFacil)

	a := 10.
	var b float64 = 3
	fmt.Println((a / b))

	var c int = 10
	d := 3
	fmt.Println((c / d))

	x := true
	y := false

	fmt.Println((x || y))
	fmt.Println((x && y))
	fmt.Println((!x))

	privada()
	Publica()
	imprimirHola()

}

func privada() {
	fmt.Println("Ejecutar logica que no necesita ser exportada")
}

//Inicia con mayuscula
func Publica() {
	fmt.Println("Logica que quiero exportar")
}

func imprimirHola() {
	defer fmt.Println("Mundo")
	fmt.Println("Hola")
}
