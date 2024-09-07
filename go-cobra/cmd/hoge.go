package cmd

import "github.com/Hayao0819/scaffold/go-cobra/cmd/hoge"

func init() {
	subCmds.Add(hoge.Cmd())
}
