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
		Dev string `envconfig:"IOTCLI_ENV_DEV",yaml:"dev"`
		Qa  string `envconfig:"IOTCLI_ENV_QA",yaml:"qa"`
	} `yaml:"environments"`
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
		return GetConfig().Envs.Qa
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
	fmt.Printf("%+v", cfg)
}
