package main

import (
	"fmt"
	"github.com/Robbie08/GoPiTest/pkg/piUtils"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
)

func main() {
	log.SetFormatter(&log.JSONFormatter{})
	log.Info("Starting server...")
    http.HandleFunc("/", defaultPage)
	http.HandleFunc("/on", onState)
	http.HandleFunc("/off", offState)
	http.HandleFunc("/shutdown", shutdown)
	http.ListenAndServe(":8080", nil)
}

func defaultPage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		log.Info("Page not found")
		return
	}
	log.Info("Someone hit the Default Page...")
	http.ServeFile(w, r, "piInterface.html")
}
func shutdown(w http.ResponseWriter, r *http.Request) {
	log.Info("Shutting server down...")
	os.Exit(0)
}
func onState(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "piInterface.html")
	fmt.Println("State: ON")
	piUtils.SendSignalToGPIO("1")
}

func offState(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "piInterface.html")
	fmt.Println("State: OFF")
	piUtils.SendSignalToGPIO("0")
}
