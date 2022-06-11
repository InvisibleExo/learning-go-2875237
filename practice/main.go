package main

import (
	"fmt"
)

func main() {
	doSomthing()

}

func doSomthing() {
	fmt.Println("Doing something from another function.")
	fmt.Println(addValues(4, 5))
	fmt.Println(addAllValues(5, 6, 6, 9, 10, 20))
}

func addValues(val1, val2 int) (sum int) {
	sum = val1 + val2
	return sum
}

func addAllValues(values ...int) (int, int) {
	total := 0
	for _, v := range values {
		total += v
	}
	return total, len(values)
}
