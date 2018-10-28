package main

import (
	"fmt"
)

// Boy a struct
type Boy struct {
	name string
	age  int
}

func (r *Boy) getBoyName() string {
	return r.name
}

// NewBoy func which creates new boy
func NewBoy(name string, age int) *Boy {
	return &Boy{name, age}
}

// PrintBoyStruct Driver of the task
func PrintBoyStruct() {
	boy := Boy{"sourav", 23}

	fmt.Printf("Contents of boy %v\n\n", boy)
	fmt.Printf("get boy name %s\n\n", boy.getBoyName())
}
