# LearningGoLang
This repository presents some examples that I used to learn GoLang.

# Examples
## Hello Word
Base code to show language structure.
> go run HelloWorld.go

## Echo
Using the "os" package to take the arguments passed to the program from the command line and concatenate them.
> go run Echo1.go Learning GoLang

## Conversor
Using the strconv package to convert a value to float64, command structure ("if" and "else") and the repetition structure ("for"). Convert KM to M and °C to K. 
> go run Conversor.go 10 20 30 c

> go run Conversor.go 5.33 12 200 km

## Calculator
Using functions and the "switch" structure to perform mathematical operations. Being used together the function "Scanf" for input of values.
> go run Calculator.go

## SameInitial
Use of package "strings" (for ToUpper function) and library (map). Returns the number of words that begin with the same letter.
> go run SameInitial.go Scissors Cuts Paper. Paper Covers Rock. Rock Crushes Lizard

## QuickSort
Aplicação do algoritmo de ordenação "QuickSort" utilizando a função append.
> go run QuickSort.go 2 65 3 85 15 48
