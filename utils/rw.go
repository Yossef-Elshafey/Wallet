package utils

import (
	"encoding/json"
	"fmt"
	"os"
	"wallet/models"
)

func GetAbsoluteFilePath() string {
	fileName := ".wallet.json"
	homePath, _ := os.UserHomeDir()
	return fmt.Sprintf("%s/%s", homePath, fileName)
}

func WriteToJsonFile(data []models.Wallet) {
	absPath := GetAbsoluteFilePath()
	json, err := json.Marshal(data)

	if err != nil {
		panic("Error while marshalling wallet data: " + err.Error())
	}

	err = os.WriteFile(absPath, json, 0644)

	if err != nil {
		panic("Error while writing to file: " + err.Error())
	}
}

func LoadJsonFile() []models.Wallet {
	absPath := GetAbsoluteFilePath()
	data, err := os.ReadFile(absPath)

	if err != nil {
		panic(err)
	}

	if len(data) == 0 {
		// At first write file is empty
		emptyArray := []models.Wallet{}
		return emptyArray
	}

	var wallets []models.Wallet
	err = json.Unmarshal(data, &wallets)

	if err != nil {
		panic("Error while unmarshalling JSON data: " + err.Error())
	}
	return wallets
}
