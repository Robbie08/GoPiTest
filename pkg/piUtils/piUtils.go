package piUtils

import (
	"fmt"
	"github.com/stianeikeland/go-rpio"
	"os"
	"time"
)

func TurnLedOn(p int) {
	fmt.Println("Pi should turn on LED soon...")
	pin := rpio.Pin(p)

	if err := rpio.Open(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer rpio.Close()

	pin.Output()

	for i := 0; i < 20; i++ {
		pin.Toggle()
		time.Sleep(time.Second)
	}
}

func TurnLedOff(p int) {
	fmt.Println("Pi should turn off LED soon...")
	pin := rpio.Pin(p)

	if err := rpio.Open(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer rpio.Close()

	pin.Output()

	for i := 0; i < 20; i++ {
		pin.Toggle()
		time.Sleep(time.Second * 4)
	}
}
