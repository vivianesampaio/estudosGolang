package main

import "fmt"

func main() {
	//ARITMÉTICOS
	soma := 10 + 10
	subtracao := 20 - 10
	divisao := 10 / 2
	multiplicacao := 10 * 2
	restoDivisao := 103 % 2

	fmt.Println(soma, subtracao, divisao, multiplicacao, restoDivisao)

	var numero = 10
	var numero2 = 30

	somaNum := numero + numero2

	fmt.Println(somaNum)

	//ATRIBUIÇÃO

	var variavel string = "String"
	teste := "Inferencia de tipos"

	fmt.Println(variavel, teste)

	//OPERADORES RELACIONAIS

	fmt.Println(1 > 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 < 2)
	fmt.Println(1 != 2)

	//OPERADORES LÓGICOS
	fmt.Println("===============")

	verdadeiro, falso := true, false
	fmt.Println(verdadeiro && falso) //&& todas as condições são verdadeiras
	fmt.Println(verdadeiro || falso) //|| uma  das condições são verdadeiras
	fmt.Println(!verdadeiro)
	fmt.Println(!falso)

	//OPERADORES UNÁRIOS
	fmt.Println("===============")

	numero10 := 10
	fmt.Println(numero10)
	numero10++
	fmt.Println(numero10)
	numero10 += 15
	fmt.Println(numero10)
	numero10 -= 10
	fmt.Println(numero10)
	numero10 *= 3
	fmt.Println(numero10)
}
