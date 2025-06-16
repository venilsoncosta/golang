package main

import "fmt"

func somar(a int, b int) int {
	return a + b
}

func retangulo(largura, altura float64) float64 {
	area := largura * altura
	return area
}

func main() {
	resultado := somar(4, 6)
	fmt.Println(resultado)

	fmt.Println("Abaixo a função retangulo")
	area := retangulo(5.0, 3.0)
	fmt.Println("Área do rentângulo:", area)
}
