package main

import "fmt"

func sum(value1 float64, value2 float64) float64{
	return value1 + value2
}

func sub(value1 float64, value2 float64) float64{
	return value1 - value2
}

func multi(value1 float64, value2 float64) float64{
	return value1 * value2
}

func div(value1 float64, value2 float64) float64{
	return float64(value1 / value2)
}

func main() {
	var i int
	var value1, value2 float64
	
	for i != -1{
		fmt.Scanf("%d %f %f", &i, &value1, &value2)
		switch i {
			case 1:
				fmt.Printf("%.f + %.f = %.f\n", value1, value2, sum(value1,value2))
			
			case 2:
				fmt.Printf("%.f - %.f = %.f\n", value1, value2, sub(value1,value2))

			case 3:
				fmt.Printf("%.f * %.f = %.f\n", value1, value2, multi(value1,value2))
			
			case 4:
				if value2 != 0 {
					fmt.Printf("%.f / %.f = %.2f\n", value1, value2, div(value1,value2))
				}else {
					fmt.Println("ERROR, IMPOSSIBLE DIVISION FOR 0")
				}
			default:
				i = -1
		}
	}
}