package main

import "fmt"

func main() {
//START OMIT
	a := make([]int, 10)
	printSlice("a", a)

	b := make([]int, 0, 5)
	printSlice("b", b)

	c := b[:2]
	printSlice("c", c)
	
	e := append(c, 10)
	printSlice("e", e)
//END OMIT
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}

