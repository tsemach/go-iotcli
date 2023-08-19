package iotinit

import (
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/tsemach/go-iotcli/cmd"
	. "github.com/tsemach/go-iotcli/cmd/common"
	"github.com/tsemach/go-iotcli/config"
	"gopkg.in/yaml.v2"
)

var iotConfigDir = First(os.UserHomeDir()) + "/.iot"

func mkdirAll(path string) {
	os.MkdirAll(path, os.ModePerm)
	i := len(path)
	color.Cyan(fmt.Sprintf("create directory %s%s|%s|", path, strings.Repeat(".", 106-i), "done"))
}

func createConfigDir() {
	// var i int

	os.Mkdir(iotConfigDir, os.ModePerm)

	envs := []string{"dev", "qa", "stage", "prod"}
	for _, e := range envs {
		mkdirAll(iotConfigDir + "/certs/" + e)
	}
	mkdirAll(iotConfigDir + "/files")
}

func downloadCert(e string, cert string) {
	var i int

	Download("http://localhost:8081/iot/certs/"+e+"/client/ca.crt", iotConfigDir+"/certs/"+e+"/"+cert)
	i = len("<" + e + ">/" + cert)
	color.Yellow(fmt.Sprintf("create certificate for %s%s|%s|", "<"+e+">/"+cert, strings.Repeat(".", 100-i), "done"))
}

func downloadCerts() {
	envs := []string{"dev", "qa", "stage", "prod"}
	// var i int
	for _, e := range envs {
		// Download("http://localhost:8081/iot/certs/"+e+"/client/ca.crt", iotConfigDir+"/certs/"+e+"/ca.crt")
		// i = len("<" + e + ">/ca.crt")
		// color.Yellow(fmt.Sprintf("create certificate for %s%s|%s|", "<"+e+">/ca.crt", strings.Repeat(".", 100-i), "done"))
		downloadCert(e, "ca.crt")
		downloadCert(e, "client.crt")
		downloadCert(e, "server.crt")

		// Download("http://localhost:8081/iot/certs/"+e+"/client/client.crt", iotConfigDir+"/certs/"+e+"/client.crt")
		// i = len("<" + e + ">/client.crt")
		// color.Yellow(fmt.Sprintf("create certificate for %s%s|%s|", "<"+e+">/client.crt", strings.Repeat(".", 100-i), "done"))

		// Download("http://localhost:8081/iot/certs/"+e+"/client/client.key", iotConfigDir+"/certs/"+e+"/client.key")
		// i = len("<" + e + ">/client.key")
		// color.Yellow(fmt.Sprintf("create certificate for %s%s|%s|", "<"+e+">/client.key", strings.Repeat(".", 100-i), "done"))
	}
}

func writeConfig() {
	var cfg config.Config
	var i int

	cfg.Dev.CAPath = iotConfigDir + "/certs/dev/ca.crt"
	cfg.Dev.ClientCrt = iotConfigDir + "/certs/dev/client.crt"
	cfg.Dev.ClientKey = iotConfigDir + "/certs/dev/client.key"

	cfg.QA.CAPath = iotConfigDir + "/certs/qa/ca.crt"
	cfg.QA.ClientCrt = iotConfigDir + "/certs/qa/client.crt"
	cfg.QA.ClientKey = iotConfigDir + "/certs/qa/client.key"

	cfg.Stage.CAPath = iotConfigDir + "/certs/stage/ca.crt"
	cfg.Stage.ClientCrt = iotConfigDir + "/certs/stage/client.crt"
	cfg.Stage.ClientKey = iotConfigDir + "/certs/stage/client.key"

	cfg.Prod.CAPath = iotConfigDir + "/certs/prod/ca.crt"
	cfg.Prod.ClientCrt = iotConfigDir + "/certs/prod/client.crt"
	cfg.Prod.ClientKey = iotConfigDir + "/certs/prod/client.key"

	yamlBytes, err := yaml.Marshal(&cfg)
	if err != nil {
		panic(err)
	}
	os.WriteFile(iotConfigDir+"/config", yamlBytes, os.ModePerm)

	i = len(iotConfigDir + "config")
	c := color.New(color.BgBlue)
	c.Print(fmt.Sprintf("create iot config file %s%s|%s|", iotConfigDir+"config", strings.Repeat(".", 100-i), "done"))
	c.DisableColor()
	fmt.Println("")
}

func writeEndMessage() {
	fmt.Println("")
	fmt.Println("iotcli configure is completed!")
	fmt.Println("")
}

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "initialize iot tool",
	Long: `this command take all the steps for initialize iot:
	1. create ~/.iot if not exist
	2. create ~/.iot/iotcli.conf, overwrite if exist
	3. create all server certificates for all envs under ~/.iot/certs/server
	4. create all client certificates for all envs under ~/.iot/certs/client`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("init called")
		createConfigDir()
		downloadCerts()
		writeConfig()
		writeEndMessage()
	},
}

func Init() {
	cmd.RootCmd.AddCommand(initCmd)
}
