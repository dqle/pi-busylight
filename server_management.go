package main

import (
	"fmt"
	"io/ioutil"
)

type serverAddress string

func getServerAddress(configFilePath string) string {
	serverAddress, _ := ioutil.ReadFile(configFilePath)
	return string(serverAddress)
}

func promptNewServerAddress() serverAddress {
	var address string
	fmt.Print("Enter Pi-Busylight Server Address: ")
	fmt.Scanln(&address)

	return serverAddress(address)
}

func (s serverAddress) saveServerAddress(configFilePath string) {
	ioutil.WriteFile(configFilePath, []byte(s), 0666)
}
