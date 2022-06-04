package main

import (
	"fmt"
)

func main() {
	poodle := Dog{"Rat Terrier", 20}
	fmt.Println(poodle)
	fmt.Printf("%+v\n", poodle)
	poodle.Weight = 25
	fmt.Printf("Breed: %+v\nWeight: %v\n", poodle.Breed, poodle.Weight)
}

// dog is a struct
type Dog struct {
	Breed  string
	Weight int
}
