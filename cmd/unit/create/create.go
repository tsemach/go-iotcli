package create

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/tsemach/go-iotcli/cmd"
	"github.com/tsemach/go-iotcli/cmd/common"
	"github.com/tsemach/go-iotcli/cmd/unit"
	"github.com/tsemach/go-iotcli/config"
)

type createStrucRequestType struct {
	Pid string `json:"pid"`
	Tid string `json:"tid"`
}

type createStructResponseType struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Route   string `json:"route"`
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "create a unit if not already exist",
	Long:  `create is part of the live cycle`,
	Run: func(_cmd *cobra.Command, args []string) {
		fmt.Println("create called")
		env := common.GetEnv(cmd.Env)

		rootCAs := common.GetRootCAs("/home/tsemach/projects/go-restapi/certs/ca.crt")
		tr := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true, RootCAs: rootCAs, ServerName: "localhost"},
		}

		body := &createStrucRequestType{
			Pid: "abc",
			Tid: "xyz",
		}
		// var body createStrucRequestType
		// body.Pid = "abc"
		// body.Tid = "xyz"

		fmt.Println("BODY:", body.Pid)
		// // m := map[string]string{"pid": "abc", "tid": "xyz"}
		r, w := io.Pipe()
		go func() {
			json.NewEncoder(w).Encode(body)
			w.Close()
		}()

		fmt.Println("URL:", fmt.Sprintf("%s%s", config.GetEnvDomain(env), "/api/v1/create"))
		client := &http.Client{Transport: tr, Timeout: 10 * time.Second}
		resp, err := client.Post(fmt.Sprintf("%s%s", config.GetEnvDomain(env), "/api/v1/create"), "application/json", r)

		// var jsonData = []byte(`{
		// 	"pid": "abc",
		// 	"tid": "xyz"
		// }`)
		// request, err := http.NewRequest("POST", fmt.Sprintf("%s/%s", config.GetEnvDomain(env), "/api/v1/create"), bytes.NewBuffer(jsonData))
		// request.Header.Set("Content-Type", "application/json; charset=UTF-8")

		// client := &http.Client{Transport: tr, Timeout: 10 * time.Second}
		// resp, err := client.Do(request)

		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()

		fmt.Println("response Status:", resp.Status)
		fmt.Println("response Headers:", resp.Header)

		var health createStructResponseType

		err = json.NewDecoder(resp.Body).Decode(&health)

		fmt.Println("create:", health)
		fmt.Println("create.status:", health.Status)
		fmt.Println("create.message:", health.Message)
		color.Cyan("Prints text in cyan.")
	},
}

func Init() {
	unit.UnitCmd.AddCommand(createCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
