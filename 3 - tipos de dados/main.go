package main

import "fmt"

func main() {
	// numeros inteiros
	numero := 100000000000
	fmt.Println(numero)

	var numero2 = -1000
	fmt.Println(numero2)
	// INT / UINT(int sem sinal)
	// alias
	// INT32 = RUNE
	var numero3 rune = 123456
	fmt.Println(numero3)

	// UINT8 = BYTE

	var numero4 byte = 78
	fmt.Println(numero4)

	// numeros reais
	//float32, float64

	var real1 float32 = 1.23
	var real2 float64 = 1.456
	var real3 = 1.99
	fmt.Println(real1, real2, real3)
}
