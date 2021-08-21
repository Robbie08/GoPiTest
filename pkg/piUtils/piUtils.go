package piUtils

import (
	"fmt"
	"os/exec"
)

func SendSignalToGPIO(state string) {
	fmt.Println("Pi should turn on LED soon...")
	LED := "7"
	gpioState := state

	// Write to GPIO pin 7
	testCmd := exec.Command("gpio", "write", LED, gpioState)
	testOut, err := testCmd.Output()

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(testOut))
	}
}
