package main

import (
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"os/user"
	"time"

	"main/icon"

	"github.com/getlantern/systray"

	"golang.org/x/sys/windows/registry"

	"github.com/gonutz/w32/v2"
)

func main() {
	systray.Run(onReady, onExit)
}

func onReady() {

	//Create Server Config File if it does not exist
	var restServerAddress serverAddress
	u, _ := user.Current()
	configFilePath := u.HomeDir + `\pi-busylight.cfg`
	_, err := os.Stat(configFilePath)
	if os.IsNotExist(err) {
		restServerAddress = promptNewServerAddress()
		restServerAddress.saveServerAddress(configFilePath)
	} else {
		restServerAddress = getServerAddress(configFilePath)
	}

	//Hide Console Windows
	console := w32.GetConsoleWindow()
	w32.ShowWindowAsync(console, w32.SW_HIDE)

	//Initiate SysTray
	systray.SetIcon(icon.Data)
	systray.SetTitle("Pi-Busylight")
	onQuit := systray.AddMenuItem("Quit", "Quit the whole app")
	onConfig := systray.AddMenuItem("Config", "Change REST Server")

	//Define where in registry we want to look for the microphone app list
	keyPath := `SOFTWARE\Microsoft\Windows\CurrentVersion\CapabilityAccessManager\ConsentStore\microphone\NonPackaged\`
	currentUser := registry.CURRENT_USER
	localMachine := registry.LOCAL_MACHINE

	//Get the list of application that uses microphone registry subkey
	currentUserSubKeyList := getMicSubKey(keyPath, currentUser)
	localMachineSubKeyList := getMicSubKey(keyPath, localMachine)

	//Infinite Loop to listen if Mic is On
	for {
		if currentUserSubKeyList.getMicOnStatus(currentUser) || localMachineSubKeyList.getMicOnStatus(localMachine) {
			_, _ = http.Post("http://"+restServerAddress.addressToString()+"/api/on", "application/json", nil)
			time.Sleep(3 * time.Second)
			systray.SetTooltip("MIC IS ON")
		} else {
			_, _ = http.Post("http://"+restServerAddress.addressToString()+"/api/off", "application/json", nil)
			time.Sleep(3 * time.Second)
			systray.SetTooltip("MIC IS OFF")
		}

		//Systray - onConfig
		//Show console again, get new server address, and Hide console
		go func() {
			<-onConfig.ClickedCh
			w32.ShowWindowAsync(console, w32.SW_SHOW)
			cmd := exec.Command("cmd", "/c", "cls")
			cmd.Stdout = os.Stdout
			cmd.Run()
			serverAddress := promptNewServerAddress()
			serverAddress.saveServerAddress(configFilePath)
			w32.ShowWindowAsync(console, w32.SW_HIDE)
		}()

		//Systray - onQuit
		go func() {
			<-onQuit.ClickedCh
			fmt.Println("Requesting quit")
			systray.Quit()
			fmt.Println("Finished quitting")
		}()
	}

}

func onExit() {
	os.Exit(1)
}
