package hoge

import (
	"os"

	"github.com/cockroachdb/errors"
	"github.com/spf13/cobra"
)

func Cmd() *cobra.Command {
	cmd := cobra.Command{
		Use:   "hoge",
		Short: "hoge command",
		RunE: func(cmd *cobra.Command, args []string) error {
			// Do something
			cmd.Println("Hello, hoge!")

			// Handle error
			pwd, err := os.Getwd()
			if err != nil {
				return errors.Wrap(err, "failed to get current directory")
			}

			cmd.Println(pwd)

			return nil
		},
	}

	return &cmd
}
