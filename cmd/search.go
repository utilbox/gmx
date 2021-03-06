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

	"github.com/spf13/cobra"
)

// searchCmd represents the search command
var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Search a module from the local collection.",
	Long:  `Search a module from the local collection.`,
	Run: func(cmd *cobra.Command, args []string) {
		searchModule()
	},
}

func init() {
	rootCmd.AddCommand(searchCmd)
}

func searchModule() {
	selected := selectModule("name of moudle to search", "module name cannot be empty")
	opts := []string{"add new version(s)",
		"use it in current project",
		"fix the module path",
		"fix a paticular version",
		"remove the module",
		"remove a paticular version",
		"exit",
	}

	idx, _ := chooseFrom("What to do?", opts)
	switch idx {
	case 0:
		addVersion(selected)
	case 1:
		useModule(selected)
	case 2:
		fixPath(selected)
	case 3:
		fixVersion(selected)
	case 4:
		removeModule(selected)
	case 5:
		removeVersion(selected)
	case 6:

	default:
		fmt.Println("Error: invalid operation")
	}

}
