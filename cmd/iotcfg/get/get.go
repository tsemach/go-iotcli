package iotcfgget

import (
	"fmt"
	"reflect"
	"strings"

	// "io"
	// "net/http"
	// "os"

	"github.com/spf13/cobra"
	"github.com/tsemach/go-iotcli/cmd/iotcfg"
	"github.com/tsemach/go-iotcli/config"
)

var key string

func examiner(t reflect.Type, depth int) {
	fmt.Println(strings.Repeat("\t", depth), "Type is", t.Name(), "and kind is", t.Kind())
	switch t.Kind() {
	case reflect.Array, reflect.Chan, reflect.Map, reflect.Ptr, reflect.Slice:
		fmt.Println(strings.Repeat("\t", depth+1), "Contained type:")
		examiner(t.Elem(), depth+1)
	case reflect.Struct:
		for i := 0; i < t.NumField(); i++ {
			f := t.Field(i)
			fmt.Println(strings.Repeat("\t", depth+1), "Field", i+1, "name is", f.Name, "type is", f.Type.Name(), "and kind is", f.Type.Kind())
			if f.Tag != "" {
				fmt.Println(strings.Repeat("\t", depth+2), "Tag is", f.Tag)
				fmt.Println(strings.Repeat("\t", depth+2), "envconfig is:", f.Tag.Get("envconfig"), "yaml tag is:", f.Tag.Get("yaml"))
			}
			if f.Type.Kind() == reflect.Ptr || f.Type.Kind() == reflect.Struct {
				examiner(f.Type, depth+1)
			}
		}
	}
}

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "get parameter from iot config file",
	Long: `iot config file can be locate with several ways:
	1. point with --config argumant
	2. point with IOTCLI_CONFIG environment variable
	3. read default from ~/.iot/config

	examples:
		1 iot config get -k dev.capath		# get certificate path of ca in dev`,
	Run: func(_cmd *cobra.Command, args []string) {
		fmt.Println("get called")
		examiner(reflect.TypeOf(config.GetConfig()), 0)

	},
}

func Init() {
	iotcfg.IOTCfgCmd.AddCommand(getCmd)

	getCmd.Flags().StringVarP(&key, "key", "k", "", "config full key with dot notation")
}
