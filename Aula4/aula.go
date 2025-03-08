package main

import "fmt"

func main() {

	var gavetas [3]string
	gavetas[0] = "comida"
	gavetas[1] = "comer"
	gavetas[2] = "quero"

	fmt.Println(gavetas[2])

	var teste []string
	teste = append(teste, "eu", "quero", "almoçar")
	fmt.Println(len(teste))
	fmt.Println(teste[:2]) //slice[x:x-1]

	var pessoas = map[string]int{}
	pessoas["lais"] = 6
	pessoas["leo"] = 26

	if idade, ok := pessoas["lais"]; ok {
		fmt.Println("Pessoas exite nesse map", idade, ok)
	} else {
		fmt.Println(pessoas)
	}

	delete(pessoas, "lais")
}
