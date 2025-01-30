package printer

import (
	"fmt"
	"reflect"
	"slices"
	"strings"
	"time"
	"wallet/models"
)

func dateDisplayFormatter(t *time.Time) string {
	return t.Format("Mon Jan 02 03:04:05 PM")
}

func getMaxColumnWidths(wallets models.Wallets, colNames []string) []int {
	vof := reflect.ValueOf(wallets[0])
	numFields := vof.NumField()
	maxColumnWidths := make([]int, 0, numFields)

	for i := 0; i < len(colNames); i++ { // head length might be wider than a value
		maxColumnWidths = append(maxColumnWidths, len(colNames[i]))
	}

	for _, wallet := range wallets {
		vof := reflect.ValueOf(wallet)
		for i := 0; i < numFields; i++ {
			fieldValue := fmt.Sprintf("%v", vof.Field(i).Interface())
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

func printHead(cn []string, columnWidth []int) {
	for gap, header := range cn {

		fmt.Printf("%-*s ", columnWidth[gap], header)
	}
	fmt.Println()

}

func printBody(wallets []models.Wallet, columnWidth []int) {
	var total int
	for _, wallet := range wallets {
		values := reflect.ValueOf(wallet)
		total += int(wallet.Amount)
		for i := 0; i < values.NumField(); i++ {
			var colValue string
			if values.Field(i).Type() == reflect.TypeOf(time.Time{}) {
				timeVal := values.Field(i).Interface().(time.Time)
				colValue = dateDisplayFormatter(&timeVal)
			} else {
				colValue = fmt.Sprintf("%v", values.Field(i).Interface())
			}

			fmt.Printf("%-*s ", columnWidth[i], colValue)
		}
		fmt.Println()
	}

	padding := slices.Max(columnWidth) / 2
	printTotal(total, padding)
}

func printTotal(total int, padding int) {
	fmt.Printf("%-*s ", padding, " ")
	fmt.Println(strings.Repeat("-", padding))
	fmt.Printf("%-*s ", padding, " ")
	fmt.Printf("Total: %d\n", total)
	fmt.Printf("%-*s ", padding, " ")
	fmt.Println(strings.Repeat("-", padding))
}

func Print(wallets []models.Wallet) {
	columnName := columnNames(wallets[0])
	maxColumnWidth := getMaxColumnWidths(wallets, columnName)
	printHead(columnName, maxColumnWidth)
	printBody(wallets, maxColumnWidth)
}
