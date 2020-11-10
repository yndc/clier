package clier

import (
	"strings"
)

// FlagConfig provides a definition of a flag
type FlagConfig struct {
	Identifier  string
	Shortcut    string
	Description string
	Handler     func()
}

// ParameterConfig provides a definition of a parameter argument
type ParameterConfig struct {
	Identifier       string
	Shortcut         string
	Description      string
	ValuePlaceholder string
	Required         bool
	Handler          func(string)
}

// PositionalConfig provides a definition of a positional argument
type PositionalConfig struct {
	Identifier  string
	Description string
	Required    bool
	Handler     func(string)
}

// Arguments hold a parsed command line arguments
type Arguments struct {
	Positional []string
	Flags      map[string]struct{}
	Parameters map[string]string
}

// HasFlag checks if any of the given flag exists in the arguments
func (a *Arguments) HasFlag(flag ...string) bool {
	for i := 0; i < len(flag); i++ {
		if _, ok := a.Flags[flag[i]]; ok {
			return true
		}
	}
	return false
}

// GetParameter retrieves the value of the given parameter identifier
func (a *Arguments) GetParameter(identifier ...string) (string, bool) {
	for i := 0; i < len(identifier); i++ {
		if val, ok := a.Parameters[identifier[i]]; ok {
			return val, true
		}
	}
	return "", false
}

// ArgType values
const (
	FLAG      = -1
	PARAMETER = -2
)

func getFlagOrParameterIdentifier(source string) (string, bool) {
	var identifier string
	if strings.HasPrefix(source, "--") {
		identifier = source[2:]
	} else if strings.HasPrefix(source, "-") {
		identifier = source[1:]
	} else {
		return identifier, false
	}
	return identifier, true
}

// Parse the given arguments and returns parsed result
func Parse(source []string) (*Arguments, error) {
	result := &Arguments{}
	result.Flags = make(map[string]struct{})
	result.Parameters = make(map[string]string)
	Process(source, func(argType int, values ...string) {
		switch argType {
		case FLAG:
			result.Flags[values[0]] = struct{}{}
		case PARAMETER:
			result.Parameters[values[0]] = values[1]
		default:
			result.Positional = append(result.Positional, values[0])
		}
	})
	return result, nil
}

// Process arguments and call the given handler
func Process(source []string, handler func(argType int, values ...string)) {
	positionalIndex := 0
	for i := 0; i < len(source); i++ {
		element := source[i]

		identifier, isFlagOrParameter := getFlagOrParameterIdentifier(element)
		if isFlagOrParameter {
			if i == (len(source) - 1) {
				handler(FLAG, identifier)
			} else {
				nextElement := source[i+1]
				_, nextIsFlagOrParameter := getFlagOrParameterIdentifier(nextElement)
				if !nextIsFlagOrParameter {
					handler(PARAMETER, identifier, nextElement)
					i++
				} else {
					handler(FLAG, identifier)
				}
			}
		} else {
			handler(positionalIndex, element)
			positionalIndex++
		}
	}
}
