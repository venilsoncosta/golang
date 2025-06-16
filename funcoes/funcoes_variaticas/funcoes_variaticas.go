package main

import "fmt"

func somarTudo(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	return total
}

func main() {
	fmt.Println("Soma de 1, 2, 3:", somarTudo(1, 2, 3))
	fmt.Println("Soma de 1, 2, 3, 4, 5:", somarTudo(1, 2, 3, 4, 5))
	fmt.Println("Soma vazia:", somarTudo())
}
