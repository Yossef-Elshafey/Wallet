package printer

import (
	"fmt"
	"reflect"
	"time"
	"wallet/models"
)

func dateDisplayFormatter(t *time.Time) string {
	return t.Format("Mon Jan 02 03:04:05 PM")
}

func getMaxColumnWidths(wallets []models.Wallet, colNames []string) []int {
	vof := reflect.ValueOf(wallets[0])
	numFields := vof.NumField()
	maxColumnWidths := make([]int, 0, numFields)

	for i := 0; i < len(colNames); i++ { // head length might be wider than a value
		maxColumnWidths = append(maxColumnWidths, len(colNames[i]))
	}

	for _, wallet := range wallets {
		vof := reflect.ValueOf(wallet)
		// fmt.Printf("%+v\n", wallet)
		for i := 0; i < numFields; i++ {
			// fmt.Printf("%v", vof.Field(i).Type() == reflect.TypeOf(time.Time{}))
			fieldValue := fmt.Sprintf("%v", vof.Field(i).Interface())
			// fmt.Printf("FieldValue: %+v", fieldValue)
			// fmt.Printf(" Value Length :%d\n", len(fieldValue))
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
	columnName := columnNames(wallets[0])
	maxColumnWidth := getMaxColumnWidths(wallets, columnName)
	for gap, header := range columnName {
		fmt.Printf("%-*s ", maxColumnWidth[gap], header)
	}

	fmt.Println()

	for _, wallet := range wallets {
		values := reflect.ValueOf(wallet)
		for i := 0; i < values.NumField(); i++ {
			var colValue string
			if values.Field(i).Type() == reflect.TypeOf(time.Time{}) {
				timeVal := values.Field(i).Interface().(time.Time)
				colValue = dateDisplayFormatter(&timeVal)
			} else {
				colValue = fmt.Sprintf("%v", values.Field(i).Interface())
			}

			fmt.Printf("%-*s ", maxColumnWidth[i], colValue)
		}
		fmt.Println()
	}
}
