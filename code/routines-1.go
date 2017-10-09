package main

import (
	"fmt"
	"time"
)

// START OMIT
func main() {
	printSaludo("Hola", 1000)
	printSaludo("Bonjour", 800)
	printSaludo("Hello", 700)
	printSaludo("Oi", 500)
}

func printSaludo(saludo string, seconds time.Duration) {
	time.Sleep((seconds * time.Millisecond))
	fmt.Println(saludo)
}

// END OMIT
