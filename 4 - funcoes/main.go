package main

import "fmt"

func somar(n1, n2 int) int {
	return n1 + n2
}

func main() {
	soma := somar(10, 20)
	fmt.Println(soma)

	var f = func(txt string) {
		fmt.Println(txt)
	}

	f("função f" + " kkkkk")
	resultadoSoma, _ := calculosMatematicos(20, 30)
	fmt.Println(resultadoSoma)
}

func calculosMatematicos(n1, n2 int) (int, int) {
	soma := n1 + n2
	subtracao := n1 - n2
	return soma, subtracao
}
