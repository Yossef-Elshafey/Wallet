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

func getMaxColumnWidths(wallets []models.Wallet) []int {
	vof := reflect.ValueOf(wallets[0])
	numFields := vof.NumField()
	maxColumnWidths := make([]int, numFields, numFields)

	for _, wallet := range wallets {
		vof := reflect.ValueOf(wallet)
		fmt.Printf("%+v\n", wallet)
		for i := 0; i < numFields; i++ {
			fieldValue := fmt.Sprintf("%v", vof.Field(i).Interface())
			// fmt.Printf("%+v\n", fieldValue)
			// fmt.Printf("Length:%d\n", len(fieldValue))
			if maxColumnWidths[i] < len(fieldValue) {
				maxColumnWidths[i] = len(fieldValue)
			}
		}
	}
	return maxColumnWidths
}

func columnNames(wallet models.Wallet) []string {
	var fieldNames []string
	vof := reflect.ValueOf(wallet)

	for i := 0; i < vof.NumField(); i++ {
		f := vof.Type().Field(i)
		fieldNames = append(fieldNames, f.Name)
	}
	return fieldNames
}

func Printer(wallets []models.Wallet) {
	maxColumnWidth := getMaxColumnWidths(wallets)
	columnName := columnNames(wallets[0])
	fmt.Println(columnName)
	fmt.Println(maxColumnWidth)
	// for _, header := range headers {
	// 	fmt.Printf("%-*s ", maxColumnWidth, header)
	// }
	// fmt.Println()
	//
	// for _, wallet := range wallets {
	// 	values := reflect.ValueOf(wallet)
	// 	for i := 0; i < values.NumField(); i++ {
	// 		colValue := fmt.Sprintf("%v", values.Field(i))
	// 		fmt.Printf("%-*s ", maxColumnWidth, colValue)
	// 	}
	// 	fmt.Println()
	// }
	// fmt.Printf("%-*s", maxColumnWidth, "#")
}
