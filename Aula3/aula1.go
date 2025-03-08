package main

import (
	"fmt"
	"strings"
)

func main() {
	var idade int8 = 2
	var idade1 int16 = 3
	var idade2 int32 = 4
	var idade3 int64 = 5

	var VAI float32 = 3.5
	var VAI1 float64 = 2.58

	var texto string = "ola"
	var texto1 string = " mundo"

	fmt.Printf("Idade (int8): %d\n", idade)
	fmt.Printf("Idade1 (int16): %d\n", idade1)
	fmt.Printf("Idade2 (int32): %d\n", idade2)
	fmt.Printf("Idade3 (int64): %d\n", idade3)

	fmt.Printf("VAI (float32): %f\n", VAI)   // Corrigido para %f
	fmt.Printf("VAI1 (float64): %f\n", VAI1) // Corrigido para %f

	fmt.Println(strings.ToUpper(texto) + strings.ToUpper(texto1))
	fmt.Println(strings.Contains(texto, "igor"))
}
