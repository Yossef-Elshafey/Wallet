package commands

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"wallet/utils"
)

type ModifyArguments struct {
	ID          int
	NewValue    float64
	NewCategory string
}

func (m *ModifyArguments) Validate() error {
	if m.NewCategory == "" && m.NewValue <= 0 {
		return fmt.Errorf("New amount value <= 0")
	}
	return nil
}

func (m *ModifyArguments) ApplyModification() {
	fileData := utils.LoadJsonFile()
	idx, err := utils.GetObjectByID(fileData, m.ID)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if m.NewValue != 0 {
		fileData[idx].Amount = m.NewValue
	}

	if m.NewCategory != "" {
		fileData[idx].Category = m.NewCategory
	}

	utils.WriteToJsonFile(fileData)
}

func RootModifyCmd() *cobra.Command {
	modifyArguments := ModifyArguments{}

	var modifyCmd = &cobra.Command{
		Use:   "modify",
		Short: "Modfiy an existing object (amount || category)",
		Run: func(cmd *cobra.Command, args []string) {
			if err := modifyArguments.Validate(); err != nil {
				fmt.Printf("Validation error: %s\n", err)
				os.Exit(1)
			}
			modifyArguments.ApplyModification()
		},
	}

	modifyCmd.Flags().IntVar(&modifyArguments.ID, "id", 0, "Object id to edit")
	modifyCmd.Flags().Float64Var(&modifyArguments.NewValue, "amount", 0, "New amount value")
	modifyCmd.Flags().StringVar(&modifyArguments.NewCategory, "category", "", "New Category value")
	modifyCmd.MarkFlagRequired("id")

	return modifyCmd
}
