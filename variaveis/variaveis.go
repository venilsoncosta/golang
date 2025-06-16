package main

import "fmt"

func main() {
	fmt.Println("Declaração curta")
	declaracaoCurta()
	fmt.Println()
	fmt.Println("Declaração longa")
	declaracaoCompleta()
	fmt.Println()
	fmt.Println("Multiplas declarações")

	multiplasDeclaracoes()
}

func declaracaoCurta() {
	sobrenome := "Viana"
	fmt.Println("Sobrenome ", sobrenome)

	quantidade := 10
	fmt.Println("Quantidade ", quantidade)

	x, y := 10, 20
	fmt.Println("x, y = ", x, y)

	sobrenome = "leite"
	fmt.Println("Sobrenome agora é ", sobrenome)

}

func declaracaoCompleta() {
	var nome string
	nome = "Louise"
	fmt.Println(nome)

	var idade int
	idade = 30
	fmt.Println(idade)

	var preco float64 = 99.99
	fmt.Println("Preço: ", preco)
}

func multiplasDeclaracoes() {
	var (
		altura      float64 = 1.75
		cidade      string  = "João Pessoa"
		ehEstudante bool
	)
	ehEstudante = true

	fmt.Println(cidade)
	fmt.Println(ehEstudante)
	fmt.Println(altura)
}
