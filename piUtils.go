package main

import (
    "fmt"
    "bufio"
    "os/exec"
)

func SendSignalToGPIO(){
    fmt.Println("Pi should turn on LED soon...")
    LED := "7"
    gpioState := 1 // 1 = ON , 0 = OFF

    // Write to GPIO pin 7
    testCmd := exec.Command("gpio","write",LED, gpioState)
    testOut, err := testCmd.Output()

    if err != nil {
        fmt.Println(err)
    }
}
