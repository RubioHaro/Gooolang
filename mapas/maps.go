package main

import "fmt"

func main() {

	m1 := make(map[string]int)
	m2 := make(map[int]string)

	m1["a"] = 8

	m2[9] = "asas"

	fmt.Println(m1)
	fmt.Println(m1["a"])

	fmt.Println(m2)
	fmt.Println(m2[9])
}
