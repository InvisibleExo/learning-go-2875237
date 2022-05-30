package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func calculateEquation(v string, val1 float64, val2 float64) (sum float64) {

	switch strings.TrimSpace(v) {
	case "+":
		sum = val1 + val2
		break
	case "-":
		sum = val1 - val2
		break
	case "*", "x":
		sum = val1 * val2
		break
	case "/":
		sum = val1 / val2
		break
	default:
		panic("You passed in an unsupported operator please in operators +, -, *, or /")
	}
	return sum
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Value 1: ")
	input1, _ := reader.ReadString('\n')
	float1, err := strconv.ParseFloat(strings.TrimSpace(input1), 64)

	if err != nil {
		fmt.Println(err)
		panic("Value 1 must be a number")
	}

	fmt.Print("Operator: ")
	inputOp, _ := reader.ReadString('\n')

	fmt.Print("Value 2: ")
	input2, _ := reader.ReadString('\n')
	float2, err := strconv.ParseFloat(strings.TrimSpace(input2), 64)
	if err != nil {
		fmt.Println(err)
		panic("Value 2 must be a number")
	}

	sum := calculateEquation(inputOp, float1, float2)
	sum = math.Round(sum*100) / 100
	fmt.Printf("The sum of %v %v %v is %v\n\n", float1, inputOp, float2, sum)
}
