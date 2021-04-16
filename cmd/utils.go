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
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"sort"
	"strings"

	"github.com/manifoldco/promptui"
	"github.com/spf13/viper"
	"github.com/utilbox/gmx/config"
)

func chooseFrom(label string, items []string) (idx int, item string) {
	s := promptui.Select{
		Label: label,
		Items: items,
	}
	idx, item, e := s.Run()
	if e != nil {
		fmt.Printf("Error: %s\n", e.Error())
		os.Exit(1)
	}
	return
}

func choosePrefixFrom(label, delimiter string, items []string) string {
	_, item := chooseFrom(label, items)
	return strings.TrimSpace(strings.Split(item, delimiter)[0])
}

func getInput(prompt, zeroErrMsg string, toLower bool) string {
	p := promptui.Prompt{
		Label: prompt,
		Validate: func(input string) error {
			input = strings.TrimSpace(input)
			if len(input) == 0 {
				return errors.New(zeroErrMsg)
			}
			return nil
		},
	}
	key, e := p.Run()
	if e != nil {
		fmt.Printf("Error: command execution failure, err: %s\n", e.Error())
		os.Exit(1)
	}
	key = strings.TrimSpace(key)
	if toLower {
		key = strings.ToLower(key)
	}
	return key
}

func removeModule(name string) {
	moduleMap := viper.AllSettings()
	delete(moduleMap, name)
	encodedConfig, _ := json.MarshalIndent(moduleMap, "", " ")
	err := viper.ReadConfig(bytes.NewReader(encodedConfig))
	if err != nil {
		fmt.Printf("Error: failed to write data, %s\n", err.Error())
		os.Exit(1)
	}
	viper.WriteConfig()
	fmt.Printf("Module %s has been successfully removed.\n", name)
}

func selectModule(prompt, zeroErrMsg string) string {
	key := getInput(prompt, zeroErrMsg, true)
	ns := []string{}

	for n, m := range config.Modules {
		if n == key {
			continue
		}
		if strings.Contains(n, key) {
			ns = append(ns, m.Name+":\t"+m.Path)
		}
	}

	sort.Slice(ns, func(i, j int) bool {
		return ns[i] < ns[j]
	})

	if m, ok := config.Modules[key]; ok {
		ns = append([]string{m.Name + ":\t" + m.Path}, ns...)
	}

	if len(ns) == 0 {
		fmt.Printf("Info: no valid module with the name %s.\n", key)
		os.Exit(1)
	}
	selected := choosePrefixFrom("choose the module", ":", ns)
	return selected
}

func listModules() string {
	ns := []string{}

	for n, m := range config.Modules {
		ns = append(ns, n+":\t"+m.Path)
	}
	if len(ns) == 0 {
		fmt.Println("Info: no valid modules in the library at the moment. Try to add one with \"gmx add\"")
		os.Exit(1)
	}
	sort.Slice(ns, func(i, j int) bool {
		return ns[i] < ns[j]
	})
	ns = append(ns, "exit")
	selected := choosePrefixFrom("choose the module", ":", ns)
	return selected
}

func runGoGet(path string) {
	c := exec.Command("go", "get", path)
	c.Stdout = os.Stdout
	c.Stdin = os.Stdin
	c.Stderr = os.Stderr
	err := c.Run()
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		os.Exit(1)
	}
}
