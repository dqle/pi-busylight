package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	scrollphathd "github.com/icco/scroll-phat-hd-go"
	"periph.io/x/periph/conn/i2c/i2creg"
	"periph.io/x/periph/host"
)

var display = initializeBus()

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "pi-busylight is running")
	//fmt.Println("Endpoint Hit: homePage")
}

func lightOn(w http.ResponseWriter, r *http.Request) {
	status := "on"
	fmt.Fprintf(w, status)
	display.Fill(0, 0, 17, 7, 100)
	display.Show()
	//fmt.Println("Endpoint Hit: lightOn")
}

func lightOff(w http.ResponseWriter, r *http.Request) {
	status := "off"
	fmt.Fprintf(w, status)
	display.Clear()
	//fmt.Println("Endpoint Hit: lightOff")
}

func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homePage)
	router.HandleFunc("/api/on", lightOn).Methods("POST")
	router.HandleFunc("/api/off", lightOff).Methods("POST")
	log.Fatal(http.ListenAndServe(":80", router))
}

func initializeBus() *scrollphathd.Display {
	_, _ = host.Init()
	bus, _ := i2creg.Open("1")
	display, _ := scrollphathd.New(bus)

	return display
}

func main() {
	fmt.Println("Starting Pi-Busylight")
	handleRequests()
}
