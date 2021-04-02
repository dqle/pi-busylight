package main

import (
	"fmt"

	"golang.org/x/sys/windows/registry"
)

type keyList []string

func openKey(constant registry.Key, keyPath string) registry.Key {
	//Set access right
	queryValue := registry.QUERY_VALUE | registry.ENUMERATE_SUB_KEYS

	//Get microphone subkey
	currentUserKey, err := registry.OpenKey(constant, keyPath, uint32(queryValue))
	if err != nil {
		fmt.Println(err)
	}

	return currentUserKey
}

func getMicSubKey(keyPath string, k registry.Key) keyList {

	//GET CURRENT_USER SUBKEY
	Key := openKey(k, keyPath)
	SubKey, err := Key.ReadSubKeyNames(-1)
	if err != nil {
		fmt.Println(err)
	}

	return keyList(joinPath(SubKey, keyPath))
}

func joinPath(subkey []string, keyPath string) []string {
	newSubkeyList := []string{}
	for _, key := range subkey {
		newSubkeyList = append(newSubkeyList, keyPath+key)
	}
	return newSubkeyList
}

func (klist keyList) getMicOnStatus(k registry.Key) bool {

	//Get LastUsedTimeStop value for CURRENT_USER or LOCAL_MACHINE
	for _, key := range klist {
		micKey := openKey(k, key)
		defer micKey.Close()
		status, _, err := micKey.GetIntegerValue("LastUsedTimeStop")
		if err != nil {
			fmt.Println(err)
		}
		if status == 0 {
			return true
		}
	}

	return false
}
