package main

import (
	"fmt"
	"os"
	"os/exec"
	"os/user"

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
	u, _ := user.Current()
	configFilePath := u.HomeDir + `\pi-busylight.cfg`
	if _, err := os.Stat(configFilePath); os.IsNotExist(err) {
		serverAddress := promptNewServerAddress()
		serverAddress.saveServerAddress(configFilePath)
	} else {
		serverAddress := getServerAddress(configFilePath)
		fmt.Println(serverAddress)
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
			systray.SetTooltip("MIC IS ON")
		} else {
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
