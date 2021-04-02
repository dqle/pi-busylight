package main

import (
	"fmt"
	"log"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "pi-busylight is running")
	fmt.Println("Endpoint Hit: homePage")
}

func lightOn(w http.ResponseWriter, r *http.Request) {
	status := "on"
	fmt.Fprintf(w, status)
	fmt.Println("Endpoint Hit: lightOn")
}

func lightOff(w http.ResponseWriter, r *http.Request) {
	status := "off"
	fmt.Fprintf(w, status)
	fmt.Println("Endpoint Hit: lightOff")
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/api/on", lightOn)
	http.HandleFunc("/api/off", lightOff)
	log.Fatal(http.ListenAndServe(":80", nil))
}

func main() {
	handleRequests()
}
