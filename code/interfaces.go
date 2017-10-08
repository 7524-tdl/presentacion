package main

import "fmt"

// START OMIT

type mascotaAnfitriona interface { // Se declara la interfaz // HL
	getSaludo() string // HL
}                   // HL
type perro struct{} // Implícitamente implementan la interfaz "mascotaAnfitriona"
type gato struct{}

func main() {
	p := perro{}
	g := gato{}
	imprimirSaludo(g)
	imprimirSaludo(p)
}
func (gato) getSaludo() string { // Implementación de la interfaz del gato // HL
	return "Mi arena. Hay que cambiarla."
}
func (perro) getSaludo() string { // Implementación de la interfaz del perro // HL
	return "Llegaste! Qué emoción! No lo puedo creer! Te amo!"
}
func imprimirSaludo(m mascotaAnfitriona) { // Función de con la interfaz como argumento // HL
	fmt.Println(m.getSaludo())
}

// END OMIT
