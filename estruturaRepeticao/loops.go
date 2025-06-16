package main

import "fmt"

func main() {
	fmt.Println("Contagem de 1 a 5: ")
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}
	fmt.Println("For como while:")

	pontos := 100
	fmt.Println("\nSimulando um jogo: ")
	for pontos > 0 {
		fmt.Printf("Você tem %d pontos. Jogando... \n", pontos)
		pontos -= 20
		if pontos <= 0 {
			fmt.Println("Game over! Seus pontos acabaram")
		}
	}

	fmt.Println("Iterando com loop fpr.. range")

	nomes := []string{"Alice", "Bob", "Charlie", "Louise", "Berth"}
	fmt.Println("Iterando sobre um slice")
	for i, nome := range nomes {
		fmt.Printf("ìndice %d: %s\n", i, nome)
	}

	idades := map[string]int{"Alice": 30, "Bob": 25, "Charlie": 35}
	fmt.Println("\nIterando sobre um map:")
	for nome, idade := range idades {
		fmt.Printf("%s tem %d anos.\n", nome, idade)
	}

	fmt.Println("\nIterando sem o índice:")
	for _, nome := range nomes {
		fmt.Println("Nome:", nome)
	}

}
