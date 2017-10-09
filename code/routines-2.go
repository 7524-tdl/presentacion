package main

import (
	"fmt"
	"time"
)

// START OMIT
func main() {
	go printSaludo("Hola", 1000)   // HL
	go printSaludo("Bonjour", 800) // HL
	go printSaludo("Hello", 700)   // HL
	go printSaludo("Oi", 500)      // HL
}

func printSaludo(saludo string, seconds time.Duration) {
	time.Sleep((seconds * time.Millisecond))
	fmt.Println(saludo)
}

// END OMIT
