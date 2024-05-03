package cmd

import "github.com/spf13/cobra"

var subCmds = []*cobra.Command{}

func Root() *cobra.Command {
	root := cobra.Command{
		Use:          "go-cobra",
		Short:        "go-cobra command",
		SilenceUsage: true,
	}

	root.AddCommand(subCmds...)
	return &root
}
