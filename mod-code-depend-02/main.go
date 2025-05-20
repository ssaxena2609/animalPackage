package main

import (
	"fmt"

	"github.com/ssaxena2609/puppy"
)

func main() {
	s1 := puppy.Bark()
	s2 := puppy.Barks()
	fmt.Println(s1)
	fmt.Println(s2)

	s3 := puppy.BigBark()
	fmt.Println(puppy.Bark())
	fmt.Println(puppy.Barks())
	fmt.Println(s3)
}
