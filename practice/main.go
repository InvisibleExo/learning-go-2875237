package main

import (
	"fmt"
	"sort"
)

func main() {
	var colors = []string{"Red", "Green", "Blue"}
	fmt.Println(colors)
	fmt.Println(colors[0])
	colors = append(colors, "Purple")
	fmt.Println(colors)

	colors = append(colors[1:])
	fmt.Println(colors)

	colors = append(colors[:len(colors)-1])
	fmt.Println(colors)

	numbers := make([]int, 5)
	numbers[0] = 134
	numbers[1] = 78
	numbers[3] = 98
	fmt.Println(numbers)

	sort.Ints(numbers)
	fmt.Println(numbers)

}
