package main

import (
	"fmt"
)

// START OMIT
func sum(x,y int){
	fmt.Println(x, "+", y ,"=",x+y)
}

func main() {
	a := 1
	b := 2
	defer sum(a,b)
	defer sum(a,a)
	defer sum(b,b)
	a = 10
	fmt.Println("a:",a,"b:",b)
}
//END OMIT
