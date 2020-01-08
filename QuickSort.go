package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	input := os.Args[1:]
	numbers := make([]int, len(input))

	for i, n := range input {
		number, _ := strconv.Atoi(n)
		numbers[i] = number
	}

	fmt.Println(input)

	fmt.Println(quicksort(numbers)) 
}

func quicksort(list []int) []int {
	if len(list) <= 1 {
		return list
	}

	n := make([]int, len(list))
	copy(n, list)

	posPivo := len(n) / 2
	pivo := n[posPivo]

	n = append(n[:posPivo], n[posPivo+1:]...)

	smallers, biggers := partition(n, pivo)

	return append(append(quicksort(smallers),pivo),quicksort(biggers)...)
}

func partition(list []int, pivo int) (smallers []int, biggers []int){
	for _, n := range list {
		if n <= pivo {
			smallers = append(smallers, n)
		}else {
			biggers = append(biggers, n)
		}
	}
	return smallers,biggers
}