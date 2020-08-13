package main

import (
	"fmt"

	"github.com/RubioHaro/calculator"
)

func main() {
	fmt.Println("Go Calculator")
	//calculator.LeerEntrada()
	input := calculator.LeerEntrada()
	operator, _ := calculator.DetectarOperador(input)
	c := calculator.Calculadora{}

	fmt.Println(c.Operar(input, operator))

}
