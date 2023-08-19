package assign

import (
	"encoding/json"
	"strconv"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/tsemach/go-iotcli/cmd"
	"github.com/tsemach/go-iotcli/cmd/common"
	"github.com/tsemach/go-iotcli/cmd/unit"
	"github.com/tsemach/go-iotcli/config"
)

var body AssignBodyStruct
var pid string

var assignCmd = &cobra.Command{
	Use:   "assign",
	Short: "unit assign api",
	Long:  `call to unit assign api as part of unit live cycle`,
	Run: func(_cmd *cobra.Command, args []string) {
		env := common.GetEnv(cmd.Env)

		err := json.Unmarshal(bodyBytes, &body)
		if err != nil {
			panic(err)
		}

		body.Pid = config.GetPid(env, pid)

		show, _ := _cmd.Flags().GetBool("show")
		if show {
			color.Cyan("Assign body")
			common.JsonPrettyPrint(body)

			return
		}

		var http = common.NewHTTP("/api/v1/assign")
		resp, ir := common.SendPost[AssignBodyStruct, AssignResponseStruct](env, http, &body)

		color.Cyan("status: " + strconv.Itoa(resp.StatusCode) + " " + ir.Status + "\nresponse: " + ir.Message)
	},
}

func Init() {
	unit.UnitCmd.AddCommand(assignCmd)
	unit.UnitCmd.AddCommand(assignCmd)

	assignCmd.Flags().StringVarP(&pid, "pid", "p", "", "unit product identifier")
	assignCmd.Flags().Bool("show", false, "print unit as a json")
}
