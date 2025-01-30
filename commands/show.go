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
	limit    int
	month    int
	category string
}

func (s *ShowArguments) Validate() error {
	if s.month > 12 || s.month < 1 {
		return fmt.Errorf("Are living on our universe, Month: %d", s.month)
	}
	return nil
}

func (s *ShowArguments) Display() {
	wallets := utils.LoadJsonFile()
	fmt.Printf("Category: %s\n", s.category)
	dataFilter, err := wallets.FilterWallet(func(wallet models.Wallet) bool {
		filterByDate := int(wallet.AddedAt.Month()) == s.month && wallet.AddedAt.Year() == time.Now().Year()
		filterByCategory := true

		if s.category != "" {
			filterByCategory = s.category == wallet.Category
		}

		return filterByDate && filterByCategory
	})

	if err != nil {
		fmt.Printf("%s: No results for Month: %d\n", err, s.month)
		os.Exit(1)
	}

	if len(dataFilter) < s.limit || s.limit <= 0 {
		s.limit = len(dataFilter)
	}

	printer.Print(dataFilter[len(dataFilter)-s.limit:])
}

func RootShowCmd() *cobra.Command {
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
	showCmd.Flags().StringVar(&showArguments.category, "category", "", "show by category")
	return showCmd
}
