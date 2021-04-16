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
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/utilbox/gmx/config"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update the info of a module in the local collection.",
	Long:  `Update the info of a module in the local collection.`,
	Run: func(cmd *cobra.Command, args []string) {
		updateModule()
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
}

func updateModule() {
	selected := selectModule("name of module to update", "module name cannot be empty")
	opts := []string{"fix module path", "fix a version"}
	i, _ := chooseFrom("what to do", opts)
	switch i {
	case 0:
		fixPath(selected)
	case 1:
		fixVersion(selected)
	default:
		fmt.Println("Error: invalid operation")
		return
	}
}

func fixPath(name string) {
	m := config.Modules[name]
	path := getInput("new path for the module", "module path cannot be empty", false)
	m.Path = path
	viper.Set(name, m)
	viper.WriteConfig()
	fmt.Printf("Path of Module %s has been successfully updated.\n", name)
}

func fixVersion(name string) {
	m := config.Modules[name]
	vs := m.Versions
	if vs == nil || len(vs) == 0 {
		fmt.Printf("Error: no valid version for Module %s.\n", name)
		os.Exit(1)
	}
	i, v := chooseFrom("choose a version to fix", vs)
	nv := getInput("new version to replace "+v, "module version cannot be empty", false)
	vs[i] = nv
	viper.Set(name, m)
	viper.WriteConfig()
	fmt.Printf("Versions of Module %s has been successfully updated.\n", name)
}
