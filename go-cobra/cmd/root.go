package cmd

import (
	"github.com/Hayao0819/nahi/cobrautils"
	"github.com/spf13/cobra"
)

var subCmds = cobrautils.Registory{}

func rootCmd() *cobra.Command {
	root := cobra.Command{
		Use:          "go-cobra",
		Short:        "go-cobra command",
		SilenceUsage: true,
		SilenceErrors: true,
	}

	subCmds.Bind(&root)
	return &root
}
