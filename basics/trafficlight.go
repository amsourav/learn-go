package main

import (
	"fmt"
)

var lightSwitchIsOn = false

// RunTraffic Run exported comment from other file
func RunTraffic() {
	printLightSwitchState()
	toogleLightSwitch()
	printLightSwitchState()
	toogleLightSwitch()
	printLightSwitchState()
}

func printLightSwitchState() {
	fmt.Println("The light switch is off: ", lightSwitchIsOn)
}

func toogleLightSwitch() {
	lightSwitchIsOn = !lightSwitchIsOn
}
