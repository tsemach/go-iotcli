package unit

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/tsemach/go-iotcli/cmd"
)

// UnitCmd represents the unit command
var UnitCmd = &cobra.Command{
	Use:   "unit",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("unit called")
	},
}

func Init() {
	cmd.RootCmd.AddCommand(UnitCmd)
}
