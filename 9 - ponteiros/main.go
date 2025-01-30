package main

import "fmt"

func main() {
	var variavel1 int = 10
	var variavel2 *int = &variavel1

	fmt.Println(variavel1, variavel2)

	variavel1++
	fmt.Println(variavel1, *variavel2)

	// // O PONTEIRO É UMA REFERÊNCIA DE MEMÓRIA
	// var variavel3 int
	// var ponteiro *int

	// variavel3 = 100
	// ponteiro = &variavel3

	// fmt.Println(variavel3, ponteiro)

	// // desreferenciação
	// variavel3 = 1005
	// fmt.Println(*ponteiro)

}
