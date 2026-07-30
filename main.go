package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {
	TirarDado(
		PedirValor("De Cuantas caras es el dado a tirar: "),
		PedirValor("Cuantas veces lo quieres que lo tire: "),
	)
}

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

// TirarDado Se emula tirar un dado y regresa el resultado de tirarlo, pide el numero de caras
// del dado a tirar
func TirarDado(caras int, repetir int) {
	for i := range repetir {

		resultado := rand.IntN(caras) + 1
		fmt.Println("El resultado de la tirana #", i+1, "es: ", resultado)
	}
}
