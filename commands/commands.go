package commands

import "github.com/spf13/cobra"

func RootCommands() []*cobra.Command {
	add := RootAddCmd()
	modify := RootModifyCmd()
	analyze := RootAnalyzeCmd()
	cmds := make([]*cobra.Command, 0)
	cmds = append(cmds, add, modify, analyze)
	return cmds
}
