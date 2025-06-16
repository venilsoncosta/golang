package main

import "fmt"

func main() {
	nota := 75

	if nota >= 60 {
		fmt.Println("Parabéns! Você foi aprovado.")
	} else {
		fmt.Println("Que pena! Você foi reprovado.")
	}

	nota = 50 // Mudando a nota para testar o else

	if nota >= 60 {
		fmt.Println("Parabéns! Você foi aprovado.")
	} else {
		fmt.Println("Que pena! Você foi reprovado.")
	}
}
