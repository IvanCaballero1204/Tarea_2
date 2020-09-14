package main

import "fmt"

func main() {
	var fahrenheit float64

	fmt.Print("Ingrese temperatura fahrenheit: ")
	fmt.Scanf("%f", &fahrenheit)

	result := (fahrenheit - 32) * 5 / 9

	fmt.Println("El valor en Celcius es:", result)
}
