package main

import "fmt"

func main() {
	var base float64
	var height float64

	fmt.Print("Ingrese valor de la base: ")
	fmt.Scanf("%f", &base)
	fmt.Print("Ingrese valor de la altura: ")
	fmt.Scanf("%f", &height)

	result := (base * height) / 2

	fmt.Println("El area del triangulo es igul a:", result)
}
