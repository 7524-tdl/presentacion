package main

import (
	"fmt"
	"time"
)

// START OMIT
func main() {
	c := make(chan string) // declaración del channel de tipo string // HL

	go printSaludo("Hola", 1000, c)
	go printSaludo("Bonjour", 800, c)
	go printSaludo("Hello", 700, c)
	go printSaludo("Oi", 500, c)

	for i := 1; i <= 4; i++ {
		fmt.Println(<-c) // de esta manera se desencolan los datos acumulados en el channel // HL
	}
}

func printSaludo(saludo string, seconds time.Duration, c chan string) {
	time.Sleep((seconds * time.Millisecond))
	c <- saludo // así se ingresan valores en el channel // HL
}

// END OMIT
