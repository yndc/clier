package clier

import (
	"fmt"
	"strings"
	"testing"
)

var cli = &CLI{
	Title:       "Test CLI",
	Description: "Some program for cli or something",
	Version:     "0.3.1",
	RootMenu: MenuNode{
		PositionalArguments: []PositionalConfig{
			{
				Identifier: "firstPos",
				Handler: func(s string) {
					fmt.Println("firstPos", s)
				},
			},
			{
				Identifier: "secondPos",
				Handler: func(s string) {
					fmt.Println("secondPos", s)
				},
			},
		},
		FlagConfigurations: []FlagConfig{
			{
				Identifier: "force",
				Shortcut:   "f",
				Handler: func() {
					fmt.Println("root force")
				},
			},
		},
		ChildMenus: map[string]MenuNode{
			"firstMenu": {
				FlagConfigurations: []FlagConfig{
					{
						Identifier: "force",
						Shortcut:   "f",
						Handler: func() {
							fmt.Println("firstMenu force")
						},
					},
				},
			},
			"secondMenu": {
				FlagConfigurations: []FlagConfig{
					{
						Identifier: "force",
						Shortcut:   "f",
						Handler: func() {
							fmt.Println("secondMenu force")
						},
					},
				},
			},
		},
	},
}

func TestMenu(t *testing.T) {
	cli.Start(strings.Split("ayy -f", " "))
}
