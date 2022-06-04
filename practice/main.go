package main

import (
	"fmt"
)

func main() {
	anInt := 42
	var p = &anInt
	fmt.Println("Value of p:", *p)

	value1 := 42.13
	pointer1 := &value1
	fmt.Println("Value fo pointer1:", *pointer1, *p)

	*pointer1 = *pointer1 / 31
	fmt.Println("Value fo pointer1:", *pointer1)
	fmt.Println("Value fo value1:", value1)
}
