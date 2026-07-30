package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {
	fmt.Println("Hola pto")
}

// Cuantas veces tirar -> logicaDado && repetir -> dar resultad
// pedir variable
// for && random
// Println

// PedirValor Ingresas un mensaje a imprimir y el usuario regresa un int
// o en caso de error se queda en un bucle hasta que ingrese un int
func PedirValor(mensaje string) int {
	var Valor int

	for {
		fmt.Print(mensaje)
		_, err := fmt.Scanln(&Valor)
		if err == nil {
			break
		}
	}

	return Valor
}

func NumeroRandom() int {
	return rand.IntN(100)
}
