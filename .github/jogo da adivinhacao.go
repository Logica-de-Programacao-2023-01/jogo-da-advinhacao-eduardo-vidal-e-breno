package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var JogarNovamente string
	tentativasTotais := 0
	var NumeroAdivinhado int
	fmt.Println("Bem vindo ao jogo da adivinhação.")

	for {
		NumeroAleatorio := rand.Intn(99) + 1
		tentativas := 0
		fmt.Print("Tente adivinhar o numero aleatorio gerado. Digite um numero inteiro entre 1 e 100: ")
		fmt.Scan(&NumeroAdivinhado)
		for {
			if NumeroAleatorio > NumeroAdivinhado {
				fmt.Println("o numero é maior que", NumeroAdivinhado)
				fmt.Print("tente novamente: ")
				fmt.Scan(&NumeroAdivinhado)
				tentativas++
				tentativasTotais++

			} else if NumeroAleatorio < NumeroAdivinhado {
				fmt.Println("o numero é menor que", NumeroAdivinhado)
				fmt.Print("tente novamente: ")
				fmt.Scan(&NumeroAdivinhado)
				tentativas++
				tentativasTotais++
			} else if NumeroAleatorio == NumeroAdivinhado {
				fmt.Println("Parabens você acertou!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!")
				tentativas++
				tentativasTotais++
				break
			}

		}
		fmt.Printf("voce fez %d tentativas", tentativas)
		fmt.Println("você deseja jogar novamente?(s/n)")
		fmt.Scan(&JogarNovamente)
		if JogarNovamente == "n" {
			break
		}
	}

}
