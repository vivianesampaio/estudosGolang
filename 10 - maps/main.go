package main

import "fmt"

func main() {
	// usuario := map[string]string{
	// 	"nome":  "viviane",
	// 	"idade": "34",
	// }

	// fmt.Println(usuairo)
	// fmt.Println(usuairo["nome"])

	usuario2 := map[string]map[string]string{
		"nome": {
			"primeiro": "viviane",
			"segundo":  "aragon",
		},
		"curso": {
			"graduação": "análise e desenvolvimento de sistemas",
		},
	}

	fmt.Println(usuario2["curso"])
	delete(usuario2, "curso")
	fmt.Println(usuario2["curso"])
	fmt.Println(usuario2)

	usuario2["signo"] = map[string]string{
		"signo": "gemeos",
	}
	fmt.Println(usuario2)

}
