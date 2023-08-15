package create

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
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

var body createStrucRequestType

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "create a unit if not already exist",
	Long:  `create is part of the live cycle`,
	Run: func(_cmd *cobra.Command, args []string) {
		env := common.GetEnv(cmd.Env)

		// isok, err := validate(&body)
		// if !isok {
		// 	fmt.Println("error on parsing:", err)
		// 	return
		// }

		rootCAs := common.GetRootCAs(config.GetCAPath(env))
		tr := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true, RootCAs: rootCAs, ServerName: "localhost"},
		}

		r, w := io.Pipe()
		go func() {
			json.NewEncoder(w).Encode(body)
			w.Close()
		}()

		client := &http.Client{Transport: tr, Timeout: 10 * time.Second}
		resp, err := client.Post(fmt.Sprintf("%s%s", config.GetEnvDomain(env), "/api/v1/create"), "application/json", r)

		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()

		var cr createStructResponseType
		err = json.NewDecoder(resp.Body).Decode(&cr)

		color.Cyan("status: " + strconv.Itoa(resp.StatusCode) + " " + cr.Status + "\nresponse: " + cr.Message)
	},
}

func Init() {
	unit.UnitCmd.AddCommand(createCmd)

	createCmd.Flags().StringVarP(&body.Pid, "pid", "p", "", "unit product identifier")
	createCmd.Flags().StringVarP(&body.Tid, "tid", "t", "", "unit tls_guid")
}

func validate(body *createStrucRequestType) (bool, error) {
	if len(body.Pid) < 15 {
		return false, fmt.Errorf("pid is less then 15")
	}

	if len(body.Pid) < 36 {
		return false, fmt.Errorf("tid is less then 15")
	}

	return true, nil
}
