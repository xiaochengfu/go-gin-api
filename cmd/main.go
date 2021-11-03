package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/xinliangnote/go-gin-api/cmd/remind"
	"os"
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
