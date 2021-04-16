/*
Copyright © 2021 Anonymous <usr_local@yeah.net>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/utilbox/gmx/config"
)

var (
	name    string
	path    string
	version string
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a module or a version to local collection.",
	Long:  `Add a module or a version to local collection.`,
	Run: func(cmd *cobra.Command, args []string) {
		add()
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}

func add() {
	name := getInput("name of the module to add", "module name cannot be empty", true)
	m, ok := config.Modules[name]
	if ok {
		fmt.Printf(`Module %s already existes:
path: %s
versions: %s
		`, name, m.Path, m.Versions)
		opts := []string{"add a new version", "exit"}
		i, _ := chooseFrom("what would you like to do?", opts)
		switch i {
		case 0:
			addVersion(name)
		case 1:

		default:
			fmt.Println("Error: invalid operation")
		}
		return
	}
	addModule(name)
}

func addModule(name string) {
	path := getInput("path of the module to add", "module path cannot be empty", false)
	m := &config.Module{Name: name, Path: path}

	opts := []string{"Yes", "No"}
	i, _ := chooseFrom("Do you want to add version(s) to the module now?", opts)
	if i == 0 {
		vs := []string{}
		filter := map[string]struct{}{}
		vs, filter = getVersWithFilter(vs, filter)
		if len(vs) > 0 {
			m.Versions = vs
		}
	}

	viper.Set(name, m)
	viper.WriteConfig()
	fmt.Printf("Info of Module %s has been successfully added.\n", name)
}

func addVersion(name string) {
	m := config.Modules[name]
	filter := map[string]struct{}{}
	vs := []string{}
	if m.Versions != nil {
		for _, v := range m.Versions {
			if _, ok := filter[v]; ok {
				continue
			}
			filter[v] = struct{}{}
			vs = append(vs, v)
		}
	}

	vs, filter = getVersWithFilter(vs, filter)
	m.Versions = vs
	viper.Set(name, m)
	viper.WriteConfig()
	fmt.Printf("Info of Module %s has been successfully updated.\n", name)
}

func getVersWithFilter(vs []string, filter map[string]struct{}) ([]string, map[string]struct{}) {
	rawVers := getInput("versions of the module to add (use ',' to delimit versions)",
		"module versions cannot be empty", false)
	versions := strings.Split(rawVers, ",")
	for _, v := range versions {
		v = strings.TrimSpace(v)
		if v == "" {
			continue
		}
		if _, ok := filter[v]; ok {
			continue
		}
		filter[v] = struct{}{}
		vs = append(vs, v)
	}
	return vs, filter
}
