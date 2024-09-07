package main

import (
	"fmt"
	"os"

	"github.com/Hayao0819/scaffold/go-cobra/cmd"
	_ "github.com/Hayao0819/scaffold/go-cobra/log"
)

func main() {
	if err := cmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "%+v\n", err)
		os.Exit(-1)
	}
}
