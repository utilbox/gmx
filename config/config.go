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
