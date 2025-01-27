package printer

import (
	"fmt"
	"reflect"
	"time"
	"wallet/models"
)

func dateDisplayFormatter(t *time.Time) string {
	return t.Format("06-Jan-02 15:04:05")
}

func getMaxRowLength(wallets []models.Wallet) int {

	maxRowLength := 0
	for _, wallet := range wallets {
		row1 := len(wallet.Category)
		row2 := len(wallet.AddedAt.String())
		requiredRowLength := max(row1, row2)
		if requiredRowLength > maxRowLength {
			maxRowLength = requiredRowLength
		}
	}
	return maxRowLength
}

func columnNames(wallet models.Wallet) []string {
	var fieldNames []string
	val := reflect.ValueOf(wallet)

	for i := 0; i < val.NumField(); i++ {
		fieldNames = append(fieldNames, val.Type().Field(i).Name)
	}
	return fieldNames
}

func NewPrinter(wallets []models.Wallet) {
	maxRowLength := getMaxRowLength(wallets) - 10

	headers := columnNames(wallets[0])
	for _, header := range headers {
		fmt.Printf("%-*s ", maxRowLength, header)
	}
	fmt.Println()

	for _, wallet := range wallets {
		values := reflect.ValueOf(wallet)
		for i := 0; i < values.NumField(); i++ {
			colValue := fmt.Sprintf("%v", values.Field(i))
			fmt.Printf("%-*s ", maxRowLength, colValue)
		}
		fmt.Println()
	}
	fmt.Printf("%-*s", maxRowLength, "#")
}
