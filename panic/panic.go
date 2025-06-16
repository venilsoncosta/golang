package main

import "fmt"

func verificaIdade(idade int) {
	if idade < 0 {
		panic("Idade não pode ser negativa")
	}
	fmt.Println("Idade válida:", idade)
}

func main() {

	fmt.Println("Início")
	verificaIdade(2)
	fmt.Println("fim")

}
