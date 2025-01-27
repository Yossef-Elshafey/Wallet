package printer

import (
	"fmt"
	"time"
	"wallet/models"
)

func DateDisplayFormatter(t *time.Time) string {
	return t.Format("06-Jan-02 15:04:05")
}

func NewPrinter(wallets []models.Wallet) {
	maxRowLength := 0
	for _, wallet := range wallets {
		row1 := len(wallet.Category)
		row2 := len(wallet.AddedAt.String())
		requiredRowLength := max(row1, row2)
		if requiredRowLength > maxRowLength {
			maxRowLength = requiredRowLength
		}
		fmt.Printf("Dealing with Wallet: %+v\n", wallet)
		fmt.Printf("Biggest length is: %d\n", maxRowLength)
	}
}
