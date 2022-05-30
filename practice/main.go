package main

import (
	"fmt"
	"math"
	"time"
)

func main() {

	n := time.Now()
	fmt.Println("I recorded this video at ", n)

	t := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), time.Now().Hour(), time.Now().Minute(), time.Now().Second(), time.Now().Nanosecond(), time.UTC)
	fmt.Println("Go launched at ", t)
	fmt.Println(t.Format(time.ANSIC))

	parsedTime, _ := time.Parse(time.ANSIC, t.String())
	fmt.Printf("The type of parsedTime is %T\n", parsedTime)

	i1, i2, i3 := 12, 45, 68
	intSum := i1 + i2 + i3

	fmt.Println("Integer Sum:", intSum)

	f1, f2, f3 := 23.5, 65.1, 76.3
	floatSum := f1 + f2 + f3
	fmt.Println("Float sum:", floatSum)

	floatSum = math.Round(floatSum*100) / 100
	fmt.Println("The sum is now, floatSum")

	circleRadius := 15.5
	circumference := circleRadius * 2 * circleRadius
	fmt.Printf("Circumference: %.2f\n", circumference)
}
