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
	Short: "add the info of a module to library",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		addModule()
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	addCmd.Flags().StringVarP(&name, "name", "n", "", "the name of the module")
	addCmd.MarkFlagRequired("name")
	addCmd.Flags().StringVarP(&path, "path", "p", "", "the full path of the module")
	addCmd.Flags().StringVarP(&version, "version", "v", "", "the version of the module")

}

func addModule() {
	if name == "" {
		fmt.Println("Error: name for the module is required")
		return
	}

	name = strings.ToLower(name)
	m, ok := config.Modules[name]
	if ok {
		if version == "" {
			fmt.Printf("Error: module %s (%s) exists, flag \"version\" is required\n", name, m.Path)
			return
		}

		for _, v := range m.Versions {
			if v != version {
				continue
			}

			fmt.Printf("Version %s exists in module %s (%s)\n", version, name, m.Path)
			return
		}

		m.Versions = append(m.Versions, version)
		viper.Set(name, m)
		viper.WriteConfig()
		return
	}

	if path == "" {
		fmt.Printf("Error: to add new module, flag \"path\" is required")
		return
	}

	m = &config.Module{Name: name, Path: path}
	if version != "" {
		m.Versions = []string{version}
	}

	viper.Set(name, m)
	viper.WriteConfig()
}
