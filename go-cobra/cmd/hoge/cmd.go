package hoge

import "github.com/spf13/cobra"

func Cmd() *cobra.Command {
	cmd := cobra.Command{
		Use:   "hoge",
		Short: "hoge command",
		RunE: func(cmd *cobra.Command, args []string) error {
			// Do something

			cmd.Println("Hello, hoge!")
			return nil
		},
	}

	return &cmd
}
