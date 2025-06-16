package main

import "fmt"

func main() {
	// Função anônima atribuída a uma variável
	saudacao := func(nome string) {
		fmt.Println("Olá,", nome)
	}
	saudacao("Mundo")

	// Função anônima executada imediatamente
	func() {
		fmt.Println("Eu sou uma função anônima executada imediatamente!")
	}()

	// Closure: a função interna "captura" a variável 'prefixo'
	geradorSaudacao := func(prefixo string) func(string) {
		return func(nome string) {
			fmt.Printf("%s, %s!\n", prefixo, nome)
		}
	}

	saudacaoPortugues := geradorSaudacao("Olá")
	saudacaoIngles := geradorSaudacao("Hello")

	saudacaoPortugues("Go")
	saudacaoIngles("Gophers")
}
