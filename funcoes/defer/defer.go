package main

import "fmt"

func minhaFuncao() {
	defer fmt.Println("Este é o terceiro a ser executado (defer)")
	defer fmt.Println("Este é o segundo a ser executado (defer)")
	fmt.Println("Este é o primeiro a ser executado")
}

func main() {
	minhaFuncao()
	fmt.Println("Fim do programa")
}
