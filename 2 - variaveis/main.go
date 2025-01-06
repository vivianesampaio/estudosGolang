package main

import "fmt"

func main() {
	var variavel string = "Variável 1"
	variavel2 := "Variável 2"
	fmt.Println(variavel)
	fmt.Println(variavel2)

	var (
		variavel3 = "variavel 3"
		variavel4 = "variavel 4"
	)

	variavel5 := "precisa do pontinho se não tiver o var"

	fmt.Println(variavel3, variavel4, variavel5)

	const constante = "Esta é uma constante"

	variavel3, variavel4 = variavel4, variavel3

	fmt.Println(variavel3, variavel4)

}
