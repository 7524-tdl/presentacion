package main

import (
	"fmt"
)

// START OMIT
func main() {
	listaSupermercado := []string{"cafe", "galletitas", "jabón"}
	mejorMateNoCafe(listaSupermercado) // HL
	fmt.Println("Tengo que comprar: ", listaSupermercado)
}

func mejorMateNoCafe(l []string) {
	l[0] = "yerba"
}

// END OMIT
