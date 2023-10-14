package iotcfgget

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"

	// "io"
	// "net/http"
	// "os"

	"github.com/spf13/cobra"
	"github.com/tsemach/go-iotcli/cmd/common"
	"github.com/tsemach/go-iotcli/cmd/iotcfg"
	"github.com/tsemach/go-iotcli/config"
)

var key string

func getTag(tag string, group string) string {
	a := strings.Split(tag, ",")

	for i := 0; i < len(a); i++ {
		parts := strings.Split(a[i], ":")

		if parts[0] == group {
			return common.First(strconv.Unquote(parts[1]))
		}
	}

	return ""
}

func findFieldByTag1(t reflect.Type, group string, tag string) (*reflect.StructField, error) {
	fmt.Println("Type is", t.Name(), "and kind is", t.Kind())

	if t.Kind() != reflect.Struct {
		return nil, fmt.Errorf("kind of %s is not struct, unable to find tag:%s:%s", t.Name(), group, tag)
	}

	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)

		fmt.Println("Field tag name is", f.Tag.Get(group), "and kind is", t.Kind())
		fmt.Println("Field", i+1, "name is", f.Name, "type is", f.Type.Name(), "and kind is", f.Type.Kind())

		fmt.Println("Field tag name is:", getTag(string(f.Tag), group), "and kind is", t.Kind())

		if getTag(string(f.Tag), group) == tag {
			fmt.Println("found tag field name is", f.Name, "and kind is", f.Type.Kind())
			return &f, nil
		}
	}

	return nil, fmt.Errorf("unable to find tag:%s:%s", group, tag)
}

func findFieldByTag(t reflect.Type, group string, tag string) (*reflect.StructField, error) {
	fmt.Println("Type is", t.Name(), "and kind is", t.Kind())

	if t.Kind() != reflect.Struct {
		return nil, fmt.Errorf("kind of %s is not struct, unable to find tag:%s:%s", t.Name(), group, tag)
	}

	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)

		fmt.Println("Field tag name is", f.Tag.Get(group), "and kind is", t.Kind())
		fmt.Println("Field", i+1, "name is", f.Name, "type is", f.Type.Name(), "and kind is", f.Type.Kind())

		fmt.Println("Field tag name is:", getTag(string(f.Tag), group), "and kind is", t.Kind())

		if getTag(string(f.Tag), group) == tag {
			fmt.Println("found tag field name is", f.Name, "and kind is", f.Type.Kind())
			return &f, nil
		}
	}

	return nil, fmt.Errorf("unable to find tag:%s:%s", group, tag)
}

func examiner(v reflect.Value, path []string, group string) (string, error) {
	t := v.Type()
	fmt.Println("[examiner] called with type:", t.Name(), ", kind is", t.Kind(), "indicator:", path)
	if t.Kind() != reflect.Struct {
		fmt.Println("[examiner] value is:", v.Elem())
	}

	switch t.Kind() {
	case reflect.Array, reflect.Chan, reflect.Map, reflect.Ptr, reflect.Slice:
		fmt.Println("found one of reflect.Array, reflect.Chan, reflect.Map, reflect.Ptr, reflect.Slice")
		return examiner(v.Elem(), path, group)
	case reflect.Struct:
		for i := 0; i < v.NumField(); i++ {
			f := v.Field(i)
			// tag := getTag(string(v.Type().Field(i).Tag), group)

			fmt.Println("[examiner] found value:", f)
			fmt.Println("[examiner] f.Type(), f.Type().Name():", f.Type(), f.Type().Name())
			fmt.Println("[examiner] f.Type().Name():", v.Type().Field(i).Tag)
			fmt.Println("[examiner] found tag", getTag(string(v.Type().Field(i).Tag), group))

			if getTag(string(v.Type().Field(i).Tag), group) == path[0] {
				fmt.Println("[examiner] found tag:", getTag(string(v.Type().Field(i).Tag), group))

				if len(path) == 1 {
					fmt.Println("founf value:", f.String())
					return f.String(), nil
				}
				return examiner(f, path[1:], group)
			}
		}
		// 	if len(path) == 1 {
		// 		f, err := findFieldByTag(t, "yaml", path[0])
		// 		if err != nil {
		// 			fmt.Println("[examiner] failed to look for field with tag:", path[0])
		// 		}

		// 		fmt.Println("[examiner] found tag", getTag(string(f.Tag), group))
		// 		fmt.Println("[examiner] found tag", reflect.ValueOf(f))
		// 		fmt.Println("[examiner] found tag", f.Type.Elem().String())

		// 		// return getTag(string(f.Tag), group), err
		// 		return getTag(string(f.Tag), group), err
		// 	}

		// 	for i := 0; i < t.NumField(); i++ {
		// 		f := t.Field(i)

		// 		fmt.Println("[examiner] found tag", f.Tag.Get(group))

		// 		if f.Tag != "" && f.Tag.Get(group) == path[0] {
		// 			fmt.Println("[examiner] found tag", f.Tag)
		// 			return examinerGetScan(f.Type, path[1:], group)
		// 		}
		// 	}

		// 	fmt.Println(fmt.Sprintf("unable to find %s:%s tag ", group, path[0]))
		// 	return "", fmt.Errorf("unable to find %s:%s tag in struct: %s", group, path[0], t.Name())
		// }
	}

	return "", nil
}

func examinerGetScan1(t reflect.Type, path []string, group string) (string, error) {
	fmt.Println("[examiner] called with type:", t.Name(), ", kind is", t.Kind(), "indicator:", path)
	// indicators := strings.Split(indicator, ".")

	switch t.Kind() {
	case reflect.Array, reflect.Chan, reflect.Map, reflect.Ptr, reflect.Slice:
		fmt.Println("found one of reflect.Array, reflect.Chan, reflect.Map, reflect.Ptr, reflect.Slice")
		return examinerGetScan(t.Elem(), path, group)
	case reflect.Struct:
		if len(path) == 1 {
			f, err := findFieldByTag(t, "yaml", path[0])
			if err != nil {
				fmt.Println("[examiner] failed to look for field with tag:", path[0])
			}

			fmt.Println("[examiner] found tag", getTag(string(f.Tag), group))
			fmt.Println("[examiner] found tag", reflect.ValueOf(f))
			fmt.Println("[examiner] found tag", f.Type.Elem().String())

			// return getTag(string(f.Tag), group), err
			return getTag(string(f.Tag), group), err
		}

		for i := 0; i < t.NumField(); i++ {
			f := t.Field(i)

			fmt.Println("[examiner] found tag", f.Tag.Get(group))

			if f.Tag != "" && f.Tag.Get(group) == path[0] {
				fmt.Println("[examiner] found tag", f.Tag)
				return examinerGetScan(f.Type, path[1:], group)
			}
		}

		fmt.Println(fmt.Sprintf("unable to find %s:%s tag ", group, path[0]))
		return "", fmt.Errorf("unable to find %s:%s tag in struct: %s", group, path[0], t.Name())
	}

	return "", fmt.Errorf("end of examinet, unable to find %s:%s tag ", group, path[0])
}

func examinerGetScan(t reflect.Type, path []string, group string) (string, error) {
	fmt.Println("[examiner] called with type:", t.Name(), ", kind is", t.Kind(), "indicator:", path)
	// indicators := strings.Split(indicator, ".")

	switch t.Kind() {
	case reflect.Array, reflect.Chan, reflect.Map, reflect.Ptr, reflect.Slice:
		fmt.Println("found one of reflect.Array, reflect.Chan, reflect.Map, reflect.Ptr, reflect.Slice")
		return examinerGetScan(t.Elem(), path, group)
	case reflect.Struct:
		if len(path) == 1 {
			f, err := findFieldByTag(t, "yaml", path[0])
			if err != nil {
				fmt.Println("[examiner] failed to look for field with tag:", path[0])
			}

			fmt.Println("[examiner] found tag", getTag(string(f.Tag), group))
			fmt.Println("[examiner] found tag", reflect.ValueOf(f))
			fmt.Println("[examiner] found tag", f.Type.Elem().String())

			// return getTag(string(f.Tag), group), err
			return getTag(string(f.Tag), group), err
		}

		for i := 0; i < t.NumField(); i++ {
			f := t.Field(i)

			fmt.Println("[examiner] found tag", f.Tag.Get(group))

			if f.Tag != "" && f.Tag.Get(group) == path[0] {
				fmt.Println("[examiner] found tag", f.Tag)
				return examinerGetScan(f.Type, path[1:], group)
			}
		}

		fmt.Println(fmt.Sprintf("unable to find %s:%s tag ", group, path[0]))
		return "", fmt.Errorf("unable to find %s:%s tag in struct: %s", group, path[0], t.Name())
	}

	return "", fmt.Errorf("end of examinet, unable to find %s:%s tag ", group, path[0])
}

func examinerOrig(t reflect.Type, depth int) {
	fmt.Println(strings.Repeat("\t", depth), "Type is", t.Name(), "and kind is", t.Kind())
	switch t.Kind() {
	case reflect.Array, reflect.Chan, reflect.Map, reflect.Ptr, reflect.Slice:
		fmt.Println(strings.Repeat("\t", depth+1), "Contained type:")
		examinerOrig(t.Elem(), depth+1)
	case reflect.Struct:
		for i := 0; i < t.NumField(); i++ {
			f := t.Field(i)

			if f.Type.Kind() == reflect.String {
				fmt.Println(strings.Repeat("\t", depth+1), "Field is type of string", depth+1, "name is", f.Name, "and kind is", f.Type.Kind())
				if f.Tag != "" {
					fmt.Println(strings.Repeat("\t", depth+2), "Tag is", f.Tag)
					fmt.Println(strings.Repeat("\t", depth+2), "envconfig is:", f.Tag.Get("envconfig"), "yaml tag is:", f.Tag.Get("yaml"))
				}
			}

			fmt.Println(strings.Repeat("\t", depth+1), "Field", depth+1, "name is", t.Name(), "and kind is", t.Kind())
			fmt.Println(strings.Repeat("\t", depth+1), "Field", i+1, "name is", f.Name, "type is", f.Type.Name(), "and kind is", f.Type.Kind())
			if f.Tag != "" {
				fmt.Println(strings.Repeat("\t", depth+2), "Tag is", f.Tag)
				fmt.Println(strings.Repeat("\t", depth+2), "envconfig is:", f.Tag.Get("envconfig"), "yaml tag is:", f.Tag.Get("yaml"))
			}

			if f.Type.Kind() == reflect.Ptr || f.Type.Kind() == reflect.Struct {
				examinerOrig(f.Type, depth+1)
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
		// examinerOrig(reflect.TypeOf(config.GetConfig()), 0)

		// value, err := examinerGetScan(reflect.TypeOf(config.GetConfig()), strings.Split(key, "."), "yaml")
		value, err := examiner(reflect.ValueOf(config.GetConfig()), strings.Split(key, "."), "yaml")
		if err != nil {
			fmt.Println("unable to examine key:", key)

			return
		}
		fmt.Println("found value:", value)
		// findFieldByTag(reflect.TypeOf(*config.GetConfig()), "yaml", "dev")
	},
}

func Init() {
	iotcfg.IOTCfgCmd.AddCommand(getCmd)

	getCmd.Flags().StringVarP(&key, "key", "k", "", "config full key with dot notation")
}
