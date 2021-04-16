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

// rmCmd represents the rm command
var rmCmd = &cobra.Command{
	Use:   "rm",
	Short: "Remove a module or a version from local collection.",
	Long:  `Remove a module or a version from local collection.`,
	Run: func(cmd *cobra.Command, args []string) {
		rmModule()
	},
}

func init() {
	rootCmd.AddCommand(rmCmd)
}

func rmModule() {
	selected := selectModule("name of module to remove", "module name cannot be empty")
	opts := []string{"remove the module", "remove a version"}
	idx, _ := chooseFrom("What to do?", opts)
	switch idx {
	case 0:
		removeModule(selected)
	case 1:
		removeVersion(selected)
	default:
		fmt.Print("Error: invalid operation")
	}

}

func removeVersion(name string) {
	m, ok := config.Modules[name]
	if !ok {
		fmt.Printf("Error: module %s doesn't exist\n", name)
		os.Exit(1)
	}
	vs := m.Versions
	if vs == nil || len(vs) == 0 {
		fmt.Printf("Error: no valid version in the info of module %s\n", name)
		os.Exit(1)
	}

	i, v := chooseFrom("choose a version to remove", vs)

	if i == 0 {
		vs = vs[1:]
	} else if i == len(vs)-1 {
		vs = vs[:i]
	} else {
		head := vs[:i]
		tail := vs[i+1:]
		vs = append(head, tail...)
	}

	m.Versions = vs
	viper.Set(name, m)
	viper.WriteConfig()
	fmt.Printf("Version %s of Module %s has been successfully removed.\n", v, name)
}
