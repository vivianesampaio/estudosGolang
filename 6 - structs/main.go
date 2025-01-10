package main

import "fmt"

type usuario struct {
	nome     string
	idade    uint
	endereco endereco
}

type endereco struct {
	endereço string
	número   int
}

func main() {
	var u usuario
	u.nome = "viviane 1"
	u.idade = 34
	fmt.Println(u)

	enderecoExemplo := endereco{endereço: "Rua dos bobos número 0"}

	u2 := usuario{nome: "viviane 2", idade: 0, endereco: enderecoExemplo}
	fmt.Println(u2)

	usuario3 := usuario{nome: "viviane 3"}
	fmt.Println(usuario3)
}
