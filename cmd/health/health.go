package health

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/tsemach/go-iotcli/cmd"
	"github.com/tsemach/go-iotcli/cmd/common"
	"github.com/tsemach/go-iotcli/config"
)

type healthStruct struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Route   string `json:"route"`
}

var healthCmd = &cobra.Command{
	Use:   "health",
	Short: "check health of a environment",
	Long: `health send https://<domain>/health API and expect 200 status code as return.
		it check the liveness of the nginx and auth service. the health of iot itself is not
		check with this API.`,
	Run: func(_cmd *cobra.Command, args []string) {
		env := common.GetEnv(cmd.Env)
		if env == "" {
			fmt.Println("ERROR: env is null, need to stop processing")

			return
		}
		fmt.Println("health called, env=" + env)

		fmt.Printf("\n[health] cfg %+s\n", fmt.Sprintf("%s/%s", config.GetEnvDomain(env), "/api"))

		rootCAs := common.GetRootCAs("/home/tsemach/projects/go-restapi/certs/ca.crt")
		cert, err := common.GetClientPair(env)

		tr := &http.Transport{
			TLSClientConfig: &tls.Config{
				ServerName:         "localhost",
				InsecureSkipVerify: true,
				Certificates:       []tls.Certificate{*cert},
				RootCAs:            rootCAs,
			},
		}

		client := &http.Client{Transport: tr, Timeout: 10 * time.Second}
		resp, err := client.Get(fmt.Sprintf("%s/%s", config.GetEnvDomain(env), "/health"))

		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()

		fmt.Println("response Status:", resp.Status)
		fmt.Println("response Headers:", resp.Header)

		var health healthStruct

		err = json.NewDecoder(resp.Body).Decode(&health)

		fmt.Println("health:", health)
		fmt.Println("health.status:", health.Status)
		fmt.Println("health.status:", health.Message)
		color.Cyan("Prints text in cyan.")

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
