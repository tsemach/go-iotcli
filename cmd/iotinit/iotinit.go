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

var iotDir = First(os.UserHomeDir()) + "/.iot"
var envs = []string{"dev", "qa", "stage", "prod"}

func mkdirAll(path string) {
	os.MkdirAll(path, os.ModePerm)
	i := len(path)
	color.Cyan(fmt.Sprintf("create directory %s%s|%s|", path, strings.Repeat(".", 106-i), "done"))
}

func createDirs() {
	os.Mkdir(iotDir, os.ModePerm)

	for _, e := range envs {
		mkdirAll(iotDir + "/certs/" + e)
	}
	mkdirAll(iotDir + "/files")
}

func downloadCert(e string, cert string) {
	var i int

	Download("http://localhost:8081/iot/certs/"+e+"/client/"+cert, iotDir+"/certs/"+e+"/"+cert)
	i = len("<" + e + ">/" + cert)
	color.Yellow(fmt.Sprintf("create certificate for %s%s|%s|", "<"+e+">/"+cert, strings.Repeat(".", 100-i), "done"))
}

func downloadCerts() {
	for _, e := range envs {
		// Download("http://localhost:8081/iot/certs/"+e+"/client/ca.crt", iotConfigDir+"/certs/"+e+"/ca.crt")
		// i = len("<" + e + ">/ca.crt")
		// color.Yellow(fmt.Sprintf("create certificate for %s%s|%s|", "<"+e+">/ca.crt", strings.Repeat(".", 100-i), "done"))
		downloadCert(e, "ca.crt")
		downloadCert(e, "admin.crt")
		downloadCert(e, "admin.key")
		downloadCert(e, "car.crt")
		downloadCert(e, "car.key")
		// downloadCert(e, "server.crt")

		// Download("http://localhost:8081/iot/certs/"+e+"/client/admin.crt", iotConfigDir+"/certs/"+e+"/admin.crt")
		// i = len("<" + e + ">/admin.crt")
		// color.Yellow(fmt.Sprintf("create certificate for %s%s|%s|", "<"+e+">/client.crt", strings.Repeat(".", 100-i), "done"))

		// Download("http://localhost:8081/iot/certs/"+e+"/client/admin.key", iotConfigDir+"/certs/"+e+"/admin.key")
		// i = len("<" + e + ">/admin.key")
		// color.Yellow(fmt.Sprintf("create certificate for %s%s|%s|", "<"+e+">/admin.key", strings.Repeat(".", 100-i), "done"))
	}
}

func writeConfig() {
	var cfg config.Config
	var i int
	var cfgPath = iotDir + "/config.yaml"

	cfg.Dev.CAPath = iotDir + "/certs/dev/ca.crt"
	cfg.Dev.AdminCrt = iotDir + "/certs/dev/admin.crt"
	cfg.Dev.AdminKey = iotDir + "/certs/dev/admin.key"
	cfg.Dev.CarCrt = iotDir + "/certs/dev/car.crt"
	cfg.Dev.CarKey = iotDir + "/certs/dev/car.key"

	cfg.QA.CAPath = iotDir + "/certs/qa/ca.crt"
	cfg.QA.AdminCrt = iotDir + "/certs/qa/admin.crt"
	cfg.QA.AdminKey = iotDir + "/certs/qa/admin.key"
	cfg.QA.CarCrt = iotDir + "/certs/qa/car.crt"
	cfg.QA.CarKey = iotDir + "/certs/qa/car.key"

	cfg.Stage.CAPath = iotDir + "/certs/stage/ca.crt"
	cfg.Stage.AdminCrt = iotDir + "/certs/stage/admin.crt"
	cfg.Stage.AdminKey = iotDir + "/certs/stage/admin.key"
	cfg.Stage.CarCrt = iotDir + "/certs/stage/car.crt"
	cfg.Stage.CarKey = iotDir + "/certs/stage/car.key"

	cfg.Prod.CAPath = iotDir + "/certs/prod/ca.crt"
	cfg.Prod.AdminCrt = iotDir + "/certs/prod/admin.crt"
	cfg.Prod.AdminKey = iotDir + "/certs/prod/admin.key"
	cfg.Prod.CarCrt = iotDir + "/certs/prod/car.crt"
	cfg.Prod.CarKey = iotDir + "/certs/prod/car.key"

	yamlBytes, err := yaml.Marshal(&cfg)
	if err != nil {
		panic(err)
	}
	os.WriteFile(cfgPath, yamlBytes, os.ModePerm)

	i = len(cfgPath)
	c := color.New(color.BgBlue)
	c.Print(fmt.Sprintf("create iot config file %s%s|%s|", cfgPath, strings.Repeat(".", 100-i), "done"))
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
	2. create ~/.iot/config.yaml, overwrite if exist	
	3. create all client certificates for all envs under ~/.iot/certs/client`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("init called")
		createDirs()
		downloadCerts()
		writeConfig()
		writeEndMessage()
	},
}

func Init() {
	cmd.RootCmd.AddCommand(initCmd)
}
