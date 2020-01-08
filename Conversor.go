package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Uso: conversor <valores> <unidade>")
		os.Exit(1)
	}

	unidadeO := os.Args[len(os.Args)-1]
	valorO := os.Args[1:len(os.Args)-1]

	var unidadeD string
	var valorD float64

	if unidadeO == "c" {
		unidadeD = "k"
	}else if unidadeO == "km" {
		unidadeD = "m"
	}else {
		fmt.Printf("%s nao e uma unidade conhecida", unidadeD)
		os.Exit(1)
	}

	for i, v := range valorO {
		valorO, err := strconv.ParseFloat(v,64)
		if err != nil {
			fmt.Println("O valor %s na posicao %d nao e um numero valido", v, i)
			os.Exit(1)
		}

		if unidadeO == "c" {
			valorD = valorO + 273
		}else{
			valorD = valorO * 1000
		}

		fmt.Printf("%.2f %s = %.2f %s\n", valorO, unidadeO, valorD, unidadeD)
	}

}