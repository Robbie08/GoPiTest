package piUtils

import (
	"bufio"
	"fmt"
	"os/exec"
)

func SendSignalToGPIO(state int) {
	fmt.Println("Pi should turn on LED soon...")
	LED := "7"
	gpioState := state

	// Write to GPIO pin 7
	testCmd := exec.Command("gpio", "write", LED, gpioState)
	testOut, err := testCmd.Output()

	if err != nil {
		fmt.Println(err)
	}
}
