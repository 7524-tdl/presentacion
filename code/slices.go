package main

import (
	"fmt"
)


func main() {
// START OMIT
	var a = [10]int{1,2,3,4,5,6,7,8,9,10,}
	fmt.Println("a:", a)
	b := a[0:5]
	c := a[3:9]
	fmt.Println("b:", b)
	fmt.Println("c:", c)
	fmt.Println("Cambio 5 por 999")
	b[4] = 999
	fmt.Println("a:", a)
	fmt.Println("b:", b)
	fmt.Println("c:", c)
	
//END OMIT
}

