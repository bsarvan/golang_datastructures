package main

import (
	"fmt"
)

// Set Data Structure
type Set struct {
	integerMap map[int]bool
}

// New method to create Set Data Structure
func (set *Set) New() {
	set.integerMap = make(map[int]bool)
}

// Contains method to check if element exists in Set
func (set *Set) Contains(element int) bool {
	_, exists := set.integerMap[element]
	return exists
}

// Add method to add an element in the Set
func (set *Set) Add(element int) {
	if !set.Contains(element) {
		set.integerMap[element] = true
	}
}

// Print method to print the contents of the Set
func (set *Set) Print() {

	fmt.Println("Contents of the Set - ")
	// for i := 0; i < 5; i++ {
	// 	fmt.Printf("%t ", set.integerMap[i])
	// }

	// Iterate a map using key and elem in the for loop
	for key, elem := range set.integerMap {
		fmt.Println(key, elem)
	}

	fmt.Printf("\n")
}

func main() {
	fmt.Println("Program to demonstrate set data structure")

	var S *Set

	S = &Set{}
	S.New()
	S.Add(1)
	S.Add(2)
	S.Add(3)

	S.Print()
	return
}
