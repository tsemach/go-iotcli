package health

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/tsemach/go-iotcli/cmd"
)

var healthCmd = &cobra.Command{
	Use:   "health",
	Short: "check health of a environment",
	Long: `health senf https://<domain>/health API and expect 200 status code as return.
		it check the liveness of the nginx and auth service. the health of iot itself is 
		check with this API.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("health called, env=" + cmd.GroupID)
	},
}

func Init() {
	cmd.RootCmd.AddCommand(healthCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// healthCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// healthCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
