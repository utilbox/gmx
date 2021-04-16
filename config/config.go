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
package config

import (
	"fmt"

	"github.com/spf13/viper"
)

var (
	Modules map[string]*Module
)

type Config struct {
	Modules []*Module
}

// Module is designed to accommodate info for a go module.
type Module struct {
	Name     string   // the last part of the path
	Path     string   // full path of the module, like "github.com/utilbox/gmx"
	Versions []string // usable versions.
}

func LoadConfig() {
	if Modules == nil {
		Modules = map[string]*Module{}
	}

	e := viper.Unmarshal(&Modules)
	if e != nil {
		fmt.Printf("failed to load .gmx.yaml: %s\n", e.Error())
		return
	}
}

func WriteConfig() {
	Modules["gin"] = &Module{
		Name: "gin",
		Path: "github.com/gin-gonic/gin",
	}
	viper.Set("", Modules)
	viper.WriteConfig()
}

// func WriteConfig() {
// 	if Modules == nil {
// 		return
// 	}

// 	ms := []*Module{}
// 	for _, m := range Modules {
// 		ms = append(ms, m)
// 	}
// 	c := &Config{Modules: ms}
// 	viper.set

// }
