package main

import (
	"fmt"
)

//START OMIT
func fact(number int) int {
	if number == 0 {
		return 1
	}
	return number * fact(number-1)
}

func main() {
	fmt.Println(fact(10))
}

//END OMIT
