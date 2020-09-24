package main

import "fmt"

func main() {
	var side float64

	fmt.Print("Ingrese el valor del lado del cuadrado: ")
	fmt.Scanf("%f", &side)

	result := side * side

	fmt.Println("El área del cuadrado es:", result)
}
