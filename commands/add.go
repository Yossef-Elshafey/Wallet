package commands

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"time"
	"wallet/models"
	"wallet/utils"
)

type AddArguments struct {
	Amount   float64
	Category string
}

type BaseValidator interface { // considered
	Validate() error
}

func (a *AddArguments) addNewWallet() []models.Wallet {
	fileData := utils.LoadJsonFile()
	wallet := models.Wallet{}

	if (len(fileData)) > 0 {
		wallet.ID = fileData[len(fileData)-1].ID + 1
	} else {
		wallet.ID = 1
	}

	wallet.Amount = a.Amount
	wallet.Category = a.Category
	wallet.AddedAt = time.Now()

	fileData = append(fileData, wallet)
	return fileData
}

func (a *AddArguments) Validate() error {
	if a.Amount <= 0 {
		return fmt.Errorf("Amount <= 0")
	}
	return nil
}

func RootAddCmd() *cobra.Command {
	addArguments := AddArguments{}

	var addCmd = &cobra.Command{
		Use:   "add",
		Short: "Create an object",
		Run: func(cmd *cobra.Command, args []string) {
			if err := addArguments.Validate(); err != nil {
				fmt.Printf("Validation error: %s\n", err)
				os.Exit(1)
			}
			newWallet := addArguments.addNewWallet()
			utils.WriteToJsonFile(newWallet)
		},
	}

	addCmd.Flags().Float64Var(&addArguments.Amount, "amount", 0, "Specify the amount value")
	addCmd.Flags().StringVar(&addArguments.Category, "category", "Not Specified", "Category name")

	return addCmd
}
