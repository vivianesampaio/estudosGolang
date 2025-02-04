package main

import "fmt"

func main() {
	fmt.Println("Estrutura de Controle - Condicionais")

	numero := -100

	if numero > 15 {
		fmt.Println("Sou maior que 15")
	} else {
		fmt.Println("Sou menor ou igual a 15")
	}

	if novaVariavel := numero; novaVariavel > 0 {
		fmt.Println(novaVariavel, "é maior que zero")
	} else if novaVariavel == -100 {
		fmt.Println(novaVariavel, "é -100")
	} else {
		fmt.Println(novaVariavel, "é menor que zero")
	}

	fmt.Println(numero, novaVariavel)
}
