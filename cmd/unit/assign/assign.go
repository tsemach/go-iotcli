/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package assign

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/tsemach/go-iotcli/cmd/unit"
)

// assignCmd represents the assign command
var assignCmd = &cobra.Command{
	Use:   "assign",
	Short: "unit assign api",
	Long:  `call to unit assign api as part of unit live cycle`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("assign called")
	},
}

func Init() {
	unit.UnitCmd.AddCommand(assignCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// assignCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// assignCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
