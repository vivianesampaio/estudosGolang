package main

import "fmt"

func main() {
	fmt.Print("Arrays")

	var array1 [1]string

	array1[0] = "viviane"

	fmt.Println(array1)

	array2 := [4]string{"1", "2", "3", "4"}
	fmt.Println(array2)
	array3 := [...]int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(array3)

	slice := []int{10, 20, 30, 40, 50, 80}
	fmt.Print("Slices")
	fmt.Println(slice)
	fmt.Println(array1)

}
