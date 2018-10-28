package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
)

// CliApp CliAPP using os arguments
func CliApp() {
	var repeatCount int
	var err error
	var gopherName string

	flag.StringVar(&gopherName, "gophername", "Gopher", "name of the gopher")
	flag.Parse()
	fmt.Println("Hello " + gopherName + " !")
	if len(os.Args) >= 2 {
		repeatCount, err = strconv.Atoi(os.Args[1])
		if err != nil {
			log.Fatal(err)
		}

		for i := 0; i < repeatCount; i++ {
			fmt.Println(os.Args[0])
			fmt.Println("Hello CLI")
		}
	} else {
		fmt.Println("Nothing Bro")
	}
}
