package main

import (
	"fmt"
	"github.com/gorilla/mux"
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
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", homePage)
	router.HandleFunc("/api/on", lightOn).Methods("PUT")
	router.HandleFunc("/api/off", lightOff).Methods("PUT")
	log.Fatal(http.ListenAndServe(":80", router))
}

func main() {
	handleRequests()
}
