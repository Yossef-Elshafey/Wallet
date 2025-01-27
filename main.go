package main

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"wallet/commands"
	"wallet/utils"
)

func createFileIfNotExist() {
	absFilePath := utils.GetAbsoluteFilePath()
	if _, err := os.Stat(absFilePath); errors.Is(err, os.ErrNotExist) {
		os.Create(absFilePath)
	}
}

func main() {
	var rootCmd = &cobra.Command{
		Use:   "wallet",
		Short: "Simple controller for colored papers(money)",
	}

	createFileIfNotExist()
	cmds := commands.RootCommands()

	for i := 0; i < len(cmds); i++ {
		rootCmd.AddCommand(cmds[i])
	}

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
