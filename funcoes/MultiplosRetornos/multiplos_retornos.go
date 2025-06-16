package main

import (
	"errors"
	"fmt"
)

func dividir(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("Divisão por zero não é permitida")
	}
	return a / b, nil
}

func main() {
	retultado, err := dividir(10, 2)
	if err != nil {
		fmt.Println("Erro:", err)
		return
	}
	fmt.Println("Resultado:", retultado)

	/*resultado, err = dividir(5, 0)
	if err != nil {
		fmt.Println("Erro:", err)
		return
	}
	fmt.Println("Resultado:", resultado)*/
}
