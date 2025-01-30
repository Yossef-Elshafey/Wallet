package models

import (
	"fmt"
	"time"
)

type Wallet struct {
	ID       int       `json:"id"`
	Amount   float64   `json:"amount"`
	Category string    `json:"category"`
	AddedAt  time.Time `json:"added_at"`
}

type Wallets []Wallet

type WalletsInterface interface {
	GetWalletByID(id int) (int, error)
	FilterWallet(func(Wallet) bool) (Wallets, error)
}

func (wallets Wallets) GetWalletByID(id int) (int, error) {
	low := 0
	high := len(wallets) - 1

	for low <= high {
		mid := (low + high) / 2

		if id == wallets[mid].ID {
			return mid, nil
		}

		if id > wallets[mid].ID {
			low = mid + 1
		} else {
			high = mid - 1
		}

	}
	return 0, fmt.Errorf("Error: Resource with id %d not found ", id)
}

func (wallets Wallets) FilterWallet(match func(Wallet) bool) (Wallets, error) {
	var result Wallets
	for _, obj := range wallets {
		if match(obj) {
			result = append(result, obj)
		}
	}

	if len(result) == 0 {
		return Wallets{}, fmt.Errorf("Not Found")
	}

	return result, nil
}
