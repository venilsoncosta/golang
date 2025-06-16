package main

import "fmt"

func calcularMedia(notas []float64) (media float64, quantidade int) {
	if len(notas) == 0 {
		return 0, 0
	}
	soma := 0.0
	for _, nota := range notas {
		soma += nota
	}
	media = soma / float64(len(notas))
	quantidade = len(notas)
	return
}

func main() {
	notas := []float64{8.5, 7.0, 9.5}
	media, quantidade := calcularMedia(notas)
	fmt.Printf("Média: %.2f, Quantidade de notas: %d\n", media, quantidade)
}
