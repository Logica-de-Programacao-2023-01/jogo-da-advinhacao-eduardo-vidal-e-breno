package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var JogarNovamente string
	tentativasTotais := []int{}
	var NumeroAdivinhado int

	fmt.Println("Bem vindo ao jogo da adivinhação. ")

	for {

		NumeroAleatorio := rand.Intn(100) + 1
		tentativas := 0
		fmt.Print("Tente adivinhar o numero aleatorio gerado. Digite um numero inteiro entre 1 e 100: ")
		fmt.Scan(&NumeroAdivinhado)
		for {
			if NumeroAleatorio > NumeroAdivinhado {
				fmt.Println("o numero é maior que", NumeroAdivinhado)
				fmt.Print("tente novamente: ")
				fmt.Scan(&NumeroAdivinhado)
				tentativas++

			} else if NumeroAleatorio < NumeroAdivinhado {
				fmt.Println("o numero é menor que", NumeroAdivinhado)
				fmt.Print("tente novamente: ")
				fmt.Scan(&NumeroAdivinhado)
				tentativas++

			} else if NumeroAleatorio == NumeroAdivinhado {
				fmt.Println("Parabens você acertou!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!")
				tentativas++
				break
			}

		}
		tentativasTotais = append(tentativasTotais, tentativas)
		fmt.Printf("voce precisou de %d chutes.", tentativas)

		fmt.Println(" deseja jogar novamente?(s/n)")
		fmt.Scan(&JogarNovamente)
		if JogarNovamente == "n" {
			for i := 0; i < len(tentativasTotais); i++ {

				fmt.Printf("voce fez %d chutes na tentativa %d . ", tentativasTotais[i], i+1)
			}
			break
		}
	}

}
