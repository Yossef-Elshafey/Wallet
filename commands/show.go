package commands

import (
	"fmt"
	"os"
	"time"
	"wallet/models"
	"wallet/printer"
	"wallet/utils"

	"github.com/spf13/cobra"
)

type ShowArguments struct {
	limit int
	month int
}

func (s *ShowArguments) Validate() error {
	if s.month > 12 || s.month < 1 {
		return fmt.Errorf("Are living on our universe, Month: %d", s.month)
	}
	return nil
}

func (s *ShowArguments) Display() {
	fileData := utils.LoadJsonFile()
	filtered := utils.FilterByMonth(fileData, func(wallet models.Wallet) bool {
		return int(wallet.AddedAt.Month()) == s.month
	})
	printer.Printer(filtered)
}

func RootAnalyzeCmd() *cobra.Command {
	showArguments := ShowArguments{}
	var showCmd = &cobra.Command{
		Use:   "show",
		Short: "Display Your colored papers outcome, Default(current month)",
		Run: func(cmd *cobra.Command, args []string) {
			if err := showArguments.Validate(); err != nil {
				fmt.Printf("Validation error: %s\n", err)
				os.Exit(1)
			}
			showArguments.Display()
		},
	}

	showCmd.Flags().IntVar(&showArguments.limit, "limit", 0, "set limitations on displayed data (loads from last)")
	showCmd.Flags().IntVar(&showArguments.month, "month", int(time.Now().Month()), "select certain month to display")
	return showCmd
}
