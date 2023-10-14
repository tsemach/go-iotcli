package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var Env string

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "go-iotcli",
	Short: "Gopher CLI in Go",
	Long:  `Gopher CLI application written in Go.`,
}

func Execute() {
	err := RootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	RootCmd.PersistentFlags().StringVarP(&Env, "env", "e", "", "select environment, one of (dev, qa, stage, prod)")
}
