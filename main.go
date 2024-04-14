package main

import (
	"fmt"

	"github.com/clydefritz/puppy"
)

func main() {
	s1 := puppy.Bark()
	s2 := puppy.Barks()
	fmt.Println(s1)
	fmt.Println(s2)

	s3 := puppy.DogBark()
	s4 := puppy.DogBarks()
	fmt.Println(s3)
	fmt.Println(s4)

	// same way result
	fmt.Println(puppy.Bark())
	fmt.Println(puppy.Barks())

}
