package main

import "fmt"

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	fmt.Println(split(17))
	x := 2
	y := 5
	g := 0.867 + 0.5i
	fmt.Println("suma", x+y)
	fmt.Printf("x is of type %T\n", x)
	fmt.Printf("g is of type %T\n", g)
}
