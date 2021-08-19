package main

import (
	"github.com/Robbie08/GoPiTest/piUtils"
	"log"
	"net/http"
	"os/exec"
)

func main() {
	http.HandleFunc("/", defaultPage)
	http.HandleFunc("/on", onState)
	http.HandleFunc("/off", offState)
	http.ListenAndServe(":8080", nil)
}

func defaultPage(w http.ResponseWritter, r *http.Request) {
	http.ServerFile(w, r, "homePage.html")
	fmt.Println("Default Page")
}

func onState(w http.ResponseWritter, r *http.Request) {
	http.ServerFile(w, r, "goButtons.html")
	fmt.Println("State: ON")
	piUtils.SendSignalToGPIO(1)
}

func offState(w http.ResponseWritter, r *http.Request) {
	http.ServerFile(w, r, "goButtons.html")
	fmt.Println("State: OFF")
	piUtils.SendSignalToGPIO(0)
}
