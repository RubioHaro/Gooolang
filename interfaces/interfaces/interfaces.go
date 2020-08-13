package main

import "fmt"

type animal interface {
	move() string
}

type dog struct{}

type fish struct{}

type bird struct{}

func (dog) move() string {
	return "perro: caminando ando"
}

func (fish) move() string {
	return "pez: nadando ando"
}

func (bird) move() string {
	return "pajaro: volando ando"
}

func moveAnimal(a animal) {
	fmt.Println(a.move())
}

func main() {
	perro := dog{}
	moveAnimal(perro)
	pez := fish{}
	moveAnimal(pez)
	pajaro := bird{}
	moveAnimal(pajaro)
}
