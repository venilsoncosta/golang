package main

import "fmt"

func main() {
	numeros := []int{10, 13, 56, 72, 91}

	fmt.Println("Iterando sobre um slice")
	for indice, valor := range numeros {
		fmt.Printf("Índice: %d, Valor: %d\n", indice, valor)
	}

	//fazendo agora sem precisar do indice
	soma := 0
	for _, valor := range numeros {
		soma += valor
	}
	fmt.Println("Soma: ", soma)
}
