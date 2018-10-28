package main

import (
	"fmt"
)

// RunMap Map Example
func RunMap() {
	var strIntMap = make(map[string]int)

	strIntMap["ON"] = 1
	strIntMap["OFF"] = 0

	fmt.Printf("COntent of map %v\n\n", strIntMap)
	fmt.Printf("Mem address of Map %p\n\n", &strIntMap)

	printAMap(strIntMap)
}

func printAMap(myMap map[string]int) {
	for key, value := range myMap {
		fmt.Printf("Value of %s is %d\n\n", key, value)
	}
}
