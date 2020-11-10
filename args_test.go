package clier

import (
	"strings"
	"testing"
)

func handleError(t *testing.T, err error) {
	if err != nil {
		t.Fatalf("%s", err)
	}
}

func handleOk(t *testing.T, ok bool) {
	if !ok {
		t.Fatalf("Not ok")
	}
}

func TestGetFlagOrParameterIdentifier(t *testing.T) {
	ans, ok := getFlagOrParameterIdentifier("-t")
	handleOk(t, ok)
	if ans != "t" {
		t.Errorf("getFlagOrParameterIdentifier(-t) = %s; want t", ans)
	}

	ans, ok = getFlagOrParameterIdentifier("--test")
	handleOk(t, ok)
	if ans != "test" {
		t.Errorf("getFlagOrParameterIdentifier(-t) = %s; want test", ans)
	}

	ans, ok = getFlagOrParameterIdentifier("test")
	if ans != "" {
		t.Errorf("getFlagOrParameterIdentifier(-t) = %s; should be empty", ans)
	}
}

func TestParse(t *testing.T) {
	ans, err := Parse(strings.Split("one two -f --default value --last", " "))
	handleError(t, err)
	t.Log(ans)
}

func TestProcess(t *testing.T) {
	params := []string{}
	positionals := []string{}
	flags := make(map[string]struct{})
	Process(strings.Split("one two three -f --aaa --par 1 -p 2 --par 3 --last", " "), func(argType int, values ...string) {
		switch argType {
		case FLAG:
			flags[values[0]] = struct{}{}
		case PARAMETER:
			params = append(params, values[1])
		default:
			positionals = append(positionals, values[0])
		}
	})

	t.Log(params)
	t.Log(positionals)
	t.Log(flags)
}
