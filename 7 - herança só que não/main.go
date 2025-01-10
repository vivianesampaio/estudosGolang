package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     int
	altura    int
}

type estudante struct {
	pessoa
	curso     string
	faculdade string
}

func main() {
	fmt.Println("Herança")
	estudante1 := pessoa{
		"viviane",
		"sampaio",
		20,
		165}
	fmt.Println(estudante1)

	estudante2 := estudante{
		estudante1,
		"analise e desenvolvimento de sistemas",
		"impacta"}
	fmt.Println(estudante2)
}
