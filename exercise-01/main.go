package main

import (
	"fmt"
)

var x = 40

const y int = 41

func main() {

	fmt.Printf("Value of x is %v and the type is: %T\n", x, x)
	fmt.Printf("Value of y is %v and the type is: %T\n", y, y)
	z := 42
	fmt.Printf("Value of z is %v and the type is: %T\n", z, z)

}
