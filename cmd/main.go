package main

import (
	"fmt"
	"os"

	"github.com/xinliangnote/go-gin-api/cmd/remind"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{}

func init() {
	rootCmd.AddCommand(remind.StartCmd)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
