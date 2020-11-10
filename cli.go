package clier

import "fmt"

// CLI is a struct managing CLI
type CLI struct {
	Title       string
	Description string
	Version     string
	RootMenu    MenuNode
}

// MenuNode defines a menu node
type MenuNode struct {
	Title                   string
	Description             string
	FlagConfigurations      []FlagConfig
	ParameterConfigurations []ParameterConfig
	PositionalArguments     []PositionalConfig
	ChildMenus              map[string]MenuNode
}

// Start the CLI application
func (c *CLI) Start(args []string) (*Arguments, error) {
	parsedArgs, err := Parse(args)
	if err != nil {
		return nil, err
	}

	node := c.RootMenu

	positionalIndexOffset := 0
	for i := 0; i < len(parsedArgs.Positional); i++ {
		if child, ok := node.ChildMenus[parsedArgs.Positional[i]]; ok {
			node = child
			positionalIndexOffset++
		} else {
			break
		}
	}

	parsedArgs.Positional = parsedArgs.Positional[positionalIndexOffset:]

	for i := 0; i < len(parsedArgs.Positional); i++ {
		if i < len(node.PositionalArguments) {
			if node.PositionalArguments[i].Handler != nil {
				node.PositionalArguments[i].Handler(parsedArgs.Positional[i])
			}
		} else {
			break
		}
	}

	for i := 0; i < len(node.FlagConfigurations); i++ {
		if _, ok := parsedArgs.Flags[node.FlagConfigurations[i].Identifier]; ok {
			if node.FlagConfigurations[i].Handler != nil {
				node.FlagConfigurations[i].Handler()
			}
		} else if _, ok := parsedArgs.Flags[node.FlagConfigurations[i].Shortcut]; ok {
			if node.FlagConfigurations[i].Handler != nil {
				node.FlagConfigurations[i].Handler()
			}

		}
	}

	for i := 0; i < len(node.ParameterConfigurations); i++ {
		if val, ok := parsedArgs.Parameters[node.ParameterConfigurations[i].Identifier]; ok {
			if node.ParameterConfigurations[i].Handler != nil {
				node.ParameterConfigurations[i].Handler(val)
			}
		} else if val, ok := parsedArgs.Parameters[node.ParameterConfigurations[i].Identifier]; ok {
			if node.ParameterConfigurations[i].Handler != nil {
				node.ParameterConfigurations[i].Handler(val)
			}
		}
	}

	return parsedArgs, nil
}

func (c *CLI) getHelp() string {
	fmt.Println(c.Title)
	fmt.Println(c.Description)
	fmt.Println("Version:", c.Version)
	return ""
}
