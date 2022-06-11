package main

import (
	"fmt"
)

func main() {
	poodle := Dog{Breed: "Terrier", Weight: 10, Sound: "Oink Oink Woof!"}
	fmt.Println(poodle)
	fmt.Printf("%+v\n", poodle)
	fmt.Printf("Breed: %v\nWeight: %v\n", poodle.Breed, poodle.Weight)

	poodle.Speak()
	poodle.SpeakThreeTimes()
	poodle.SpeakThreeTimesFromMemory()
	poodle.SpeakThreeTimesFromMemory()
	poodle.SpeakThreeTimesFromMemory()
}

// Dog is a struct
type Dog struct {
	Breed  string
	Weight int
	Sound  string
}

// Speak is how the dog speaks
func (d Dog) Speak() {
	fmt.Println(d.Sound)
}

// Lets the dog speak 3 times
func (d Dog) SpeakThreeTimes() {
	d.Sound = fmt.Sprintf("1: %v :: 2: %v :: 3: %v ::", d.Sound, d.Sound, d.Sound)
	fmt.Println(d.Sound)
}

// Lets the dog speak 3 times and what it remembers from before
func (d *Dog) SpeakThreeTimesFromMemory() {
	d.Sound = fmt.Sprintf("1: %v :: 2:  %v :: 3: %v ::", d.Sound, d.Sound, d.Sound)
	fmt.Println(d.Sound)
}
