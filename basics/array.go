package main

import (
	"fmt"
)

func populateArray(intArray [10]int) {
	for i := 0; i < len(intArray)-1; i++ {
		intArray[i] = 10 * i
	}
}

func populateArrayAndReturn(intArray [10]int) [10]int {
	for i := 0; i < len(intArray)-1; i++ {
		intArray[i] = 10 * i
	}

	return intArray
}

// RunArray runs the array programs
func RunArray() {
	var integerArray [10]int
	var stringArray [5]string

	integerArray[2] = 10
	populateArray(integerArray)
	fmt.Printf("Contents of my array %v\n\n", integerArray)
	fmt.Printf("Contents of my array %v\n\n", stringArray)
	fmt.Printf("Mem address for intArray %p\n\n", &integerArray)
	integerArray = populateArrayAndReturn(integerArray)
	fmt.Printf("Contents of my array %v\n\n", integerArray)
	fmt.Printf("Mem address for intArray %p\n\n", &integerArray)
	fmt.Printf("Mem address for strArray %p\n\n", &stringArray)

}
