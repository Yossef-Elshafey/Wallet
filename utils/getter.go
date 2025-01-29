package utils

import (
	"fmt"
	"wallet/models"
)

func GetObjectByID(data []models.Wallet, id int) (int, error) {
	fileData := LoadJsonFile()
	low := 0
	high := len(fileData) - 1

	for low <= high {
		mid := (low + high) / 2

		if id == fileData[mid].ID {
			return mid, nil
		}

		if id > fileData[mid].ID {
			low = mid + 1
		} else {
			high = mid - 1
		}

	}

	return 0, fmt.Errorf("Error: Resource with id %d not found ", id)
}

func FilterByMonth(data []models.Wallet, targetMonth func(models.Wallet) bool) ([]models.Wallet, error) {
	var result []models.Wallet
	for _, obj := range data {
		if targetMonth(obj) {
			result = append(result, obj)
		}
	}
	if len(result) == 0 {
		return []models.Wallet{}, fmt.Errorf("Not Found")
	}
	return result, nil
}
