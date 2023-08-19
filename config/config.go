package config

// from: https://dev.to/ilyakaznacheev/a-clean-way-to-pass-configs-in-a-go-application-1g64

import (
	"fmt"
	"os"

	"github.com/kelseyhightower/envconfig"
	"gopkg.in/yaml.v2"
)

type Config struct {
	Envs struct {
		Dev   string `envconfig:"IOTCLI_ENV_DEV",yaml:"dev"`
		QA    string `envconfig:"IOTCLI_ENV_QA",yaml:"qa"`
		Stage string `envconfig:"IOTCLI_ENV_QA",yaml:"stage"`
		Prod  string `envconfig:"IOTCLI_ENV_QA",yaml:"prod"`
	} `yaml:"environments"`
	Dev struct {
		CAPath   string `envconfig:"IOTCLI_CAPATH",yaml:"capath"`
		AdminKey string `envconfig:"IOTCLI_ADMIN_KEY",yaml:"adminkey"`
		AdminCrt string `envconfig:"IOTCLI_ADMIN_CRT",yaml:"admincrt"`
		CarKey   string `envconfig:"IOTCLI_CLIENT_KEY",yaml:"carkey"`
		CarCrt   string `envconfig:"IOTCLI_CLIENT_CRT",yaml:"carcrt"`
		Pid      string `envconfig:"IOTCLI_PID",yaml:"pid"`
		Tid      string `envconfig:"IOTCLI_TID",yaml:"tid"`
	} `yaml:"dev"`
	QA struct {
		CAPath   string `envconfig:"IOTCLI_CAPATH",yaml:"capath"`
		AdminKey string `envconfig:"IOTCLI_ADMIN_KEY",yaml:"adminkey"`
		AdminCrt string `envconfig:"IOTCLI_ADMIN_CRT",yaml:"admincrt"`
		CarKey   string `envconfig:"IOTCLI_CLIENT_KEY",yaml:"carkey"`
		CarCrt   string `envconfig:"IOTCLI_CLIENT_CRT",yaml:"carcrt"`
		Pid      string `envconfig:"IOTCLI_PID",yaml:"pid"`
		Tid      string `envconfig:"IOTCLI_TID",yaml:"tid"`
	} `yaml:"qa"`
	Stage struct {
		CAPath   string `envconfig:"IOTCLI_CAPATH",yaml:"capath"`
		AdminKey string `envconfig:"IOTCLI_ADMIN_KEY",yaml:"adminkey"`
		AdminCrt string `envconfig:"IOTCLI_ADMIN_CRT",yaml:"admincrt"`
		CarKey   string `envconfig:"IOTCLI_CLIENT_KEY",yaml:"carkey"`
		CarCrt   string `envconfig:"IOTCLI_CLIENT_CRT",yaml:"carcrt"`
		Pid      string `envconfig:"IOTCLI_PID",yaml:"pid"`
		Tid      string `envconfig:"IOTCLI_TID",yaml:"tid"`
	} `yaml:"stage"`
	Prod struct {
		CAPath   string `envconfig:"IOTCLI_CAPATH",yaml:"capath"`
		AdminKey string `envconfig:"IOTCLI_ADMIN_KEY",yaml:"adminkey"`
		AdminCrt string `envconfig:"IOTCLI_ADMIN_CRT",yaml:"admincrt"`
		CarKey   string `envconfig:"IOTCLI_CLIENT_KEY",yaml:"carkey"`
		CarCrt   string `envconfig:"IOTCLI_CLIENT_CRT",yaml:"carcrt"`
		Pid      string `envconfig:"IOTCLI_PID",yaml:"pid"`
		Tid      string `envconfig:"IOTCLI_TID",yaml:"tid"`
	} `yaml:"prod"`
}

var cfg Config

func GetConfig() *Config {
	return &cfg
}

func GetEnvDomain(env string) string {
	switch env {
	case "dev":
		return GetConfig().Envs.Dev
	case "qa":
		return GetConfig().Envs.QA
	}

	return ""
}

func GetCAPath(env string) string {
	switch env {
	case "dev":
		return GetConfig().Dev.CAPath
	case "qa":
		return GetConfig().QA.CAPath
	}

	return ""
}

func GetAdminCert(env string) (string, string) {
	switch env {
	case "dev":
		return GetConfig().Dev.AdminCrt, GetConfig().Dev.AdminKey
	case "qa":
		return GetConfig().QA.AdminCrt, GetConfig().QA.AdminKey
	case "stage":
		return GetConfig().Stage.AdminCrt, GetConfig().Stage.AdminKey
	case "prod":
		return GetConfig().Prod.AdminCrt, GetConfig().Prod.AdminKey
	}

	return "", ""
}

func GetPid(env string, pid string) string {
	if pid != "" {
		return pid
	}

	switch env {
	case "dev":
		return GetConfig().Dev.Pid
	case "qa":
		return GetConfig().QA.Pid
	}

	return ""
}

func GetTid(env string, tid string) string {
	if tid != "" {
		return tid
	}

	switch env {
	case "dev":
		return GetConfig().Dev.Tid
	case "qa":
		return GetConfig().QA.Tid
	}

	return ""
}

func processError(err error) {
	fmt.Println(err)
}

func readFile(cfg *Config) {
	f, err := os.Open("iotcli.yml")
	if err != nil {
		processError(err)
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(cfg)
	if err != nil {
		processError(err)
	}
}

func readEnv(cfg *Config) {
	err := envconfig.Process("", cfg)
	if err != nil {
		processError(err)
	}
}

func BuildConfig() {
	readFile(&cfg)
	readEnv(&cfg)
}
