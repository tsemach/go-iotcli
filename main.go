package main

import (
	"github.com/tsemach/go-iotcli/config"

	"github.com/fatih/color"
	"github.com/tsemach/go-iotcli/cmd"
	"github.com/tsemach/go-iotcli/cmd/health"
	"github.com/tsemach/go-iotcli/cmd/unit"
	"github.com/tsemach/go-iotcli/cmd/unit/assign"
	"github.com/tsemach/go-iotcli/cmd/unit/create"
	"github.com/tsemach/go-iotcli/cmd/unit/install"
	"github.com/tsemach/go-iotcli/cmd/update"
)

func main() {
	config.BuildConfig()
	color.Yellow("using url: " + config.GetEnvDomain("dev"))

	health.Init()
	unit.Init()
	create.Init()
	assign.Init()
	install.Init()
	update.Init()

	cmd.Execute()
}
