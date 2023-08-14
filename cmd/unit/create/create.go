package create

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
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

		fmt.Println("ENV:", env)
		rootCAs := common.GetRootCAs("/home/tsemach/projects/go-restapi/certs/ca.crt")
		tr := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true, RootCAs: rootCAs, ServerName: "localhost"},
		}

		body := &createStrucRequestType{
			Pid: "abc",
			Tid: "xyz",
		}
		postBody, _ := json.Marshal(body)
		responseBody := bytes.NewBuffer(postBody)

		client := &http.Client{Transport: tr, Timeout: 10 * time.Second}
		fmt.Println("URL:" + fmt.Sprintf("%s/%s", config.GetEnvDomain(env), "/api/v1/create"))
		resp, err := client.Post(fmt.Sprintf("%s%s", config.GetEnvDomain(env), "/api/v1/create"), "application/json", responseBody)
		// resp, err := client.Post("https://localhost:8080/api/v1/create", "application/json", responseBody)

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
