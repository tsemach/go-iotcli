package install

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/tsemach/go-iotcli/cmd"
	"github.com/tsemach/go-iotcli/cmd/common"
	"github.com/tsemach/go-iotcli/cmd/unit"
	"github.com/tsemach/go-iotcli/config"
)

var body InstallBodyStruct
var pid string
var tid string

var installCmd = &cobra.Command{
	Use:   "install",
	Short: "install unit",
	Long:  `send http to install unit on iot server`,
	Run: func(_cmd *cobra.Command, args []string) {
		env := common.GetEnv(cmd.Env)

		err := json.Unmarshal(bodyBytes, &body)
		if err != nil {
			panic(err)
		}

		body.Pid = config.GetPid(env, pid)
		body.Tid = config.GetTid(env, tid)

		show, _ := _cmd.Flags().GetBool("show")
		if show {
			color.Cyan("Install body")
			common.JsonPrettyPrint(body)

			return
		}

		client := common.GetClient(env)

		postBody, _ := json.Marshal(body)
		responseBody := bytes.NewBuffer(postBody)
		resp, err := client.Post(fmt.Sprintf("%s%s", config.GetEnvDomain(env), "/api/v1/install"), "application/json", responseBody)

		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()

		var ir installStructResponseType
		err = json.NewDecoder(resp.Body).Decode(&ir)

		color.Cyan("status: " + strconv.Itoa(resp.StatusCode) + " " + ir.Status + "\nresponse: " + ir.Message)
	},
}

func Init() {
	unit.UnitCmd.AddCommand(installCmd)

	installCmd.Flags().StringVarP(&pid, "pid", "p", "", "unit product identifier")
	installCmd.Flags().StringVarP(&tid, "tid", "t", "", "unit tls guid")
	installCmd.Flags().Bool("show", false, "print unit as a json")
}
