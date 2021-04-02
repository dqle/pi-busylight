package main

import (
	"fmt"
	"log"
	"net/http"
)

func getLightStatus(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "true")
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
	http.HandleFunc("/", getLightStatus)
	http.HandleFunc("/api/on", lightOn)
	http.HandleFunc("/api/off", lightOff)
	log.Fatal(http.ListenAndServe(":10000", nil))
}

func main() {
	handleRequests()
}
