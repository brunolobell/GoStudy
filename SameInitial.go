package main

import (
	"fmt"
	"os"
	"strings"
)

func pickStatistics(words []string) map[string]int {
	statistics := make(map[string]int)
	
	
	for _, word := range words {
		initial := strings.ToUpper(string(word[0]))
		count, found := statistics[initial]		 		
		if found {
			statistics[initial] = count + 1
		}else {
			statistics[initial] = 1
		}
	}
	return statistics
}

func show(statistics map[string]int) {
	for initial, i := range statistics {
		fmt.Printf("%s = %d\n",initial, i)
	}
}

func main() {
	words := os.Args[1:]
	statistics := pickStatistics(words)
	show(statistics)
}