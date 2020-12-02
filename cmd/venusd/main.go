package main

import (
	"os"

	"github.com/fadeev/venus/cmd/venusd/cmd"
)

func main() {
	rootCmd, _ := cmd.NewRootCmd()
	if err := cmd.Execute(rootCmd); err != nil {
		os.Exit(1)
	}
}
