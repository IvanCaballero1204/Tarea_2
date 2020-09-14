package main

import "fmt"

func main() {
	const pi float64 = 3.1416
	var radius float64

	fmt.Print("Ingrese radio del circulo: ")
	fmt.Scanf("%f", &radius)

	result := pi * (radius * radius)

	fmt.Println("El area del circulo es:", result)
}
