package iotcfg

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/tsemach/go-iotcli/cmd"
)

var IOTCfgCmd = &cobra.Command{
	Use:   "config",
	Short: "set / get configuration file parameters",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("iotcfg called")
	},
}

func Init() {
	cmd.RootCmd.AddCommand(IOTCfgCmd)
}
