package main

import "fmt"

type dog struct{}

type fish struct{}

type bird struct{}

func (dog) walk() string {
	return "perro: caminando ando"
}

func (fish) swim() string {
	return "pez: nadando ando"
}

func (bird) fly() string {
	return "pajaro: volando ando"
}

func moveDog(d dog) {
	fmt.Println(d.walk())
}

func moveFish(f fish) {
	fmt.Println(f.swim())
}

func moveBird(b bird) {
	fmt.Println(b.fly())
}

func main() {
	perro := dog{}
	moveDog(perro)
	pez := fish{}
	moveFish(pez)
	pajaro := bird{}
	moveBird(pajaro)
}
