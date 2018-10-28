package main

import (
	"flag"
	"fmt"
	"log"
	"regexp"
)

// UsernameRegex const regexmfor username checking
const UsernameRegex string = `^@?(\w){1,15}$`

// RunUsernameChecker runs the username validator
func RunUsernameChecker() {
	var usernameFlag string
	flag.StringVar(&usernameFlag, "username", "@sourav", "username of the user")
	flag.Parse()

	if validaterUsername(usernameFlag) {
		fmt.Println("you have entered a VALID username")
		RunTraffic()
	} else {
		fmt.Println("you have entered a INVALID username")
	}

}

func validaterUsername(username string) bool {
	usernameRegexMatcher, regexErr := regexp.Compile(UsernameRegex)

	if regexErr != nil {
		log.Fatal("something went wrong")
	} else {
		return usernameRegexMatcher.MatchString(username)
	}
	return false
}
