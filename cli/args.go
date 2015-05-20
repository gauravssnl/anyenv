package cli

import (
	"path"
	"path/filepath"
	"strings"
)

type Args struct {
	ARGV []string
}

func (a Args) ProgramName() string {
	return path.Base(a.ARGV[0])
}

func (a Args) FullProgramName() string {
	abs, err := filepath.Abs(a.ARGV[0])
	if err == nil {
		return abs
	} else {
		return a.ARGV[0]
	}
}

func (a Args) CommandName() string {
	return a.At(-1)
}

func (a Args) At(n int) string {
	n += 2
	if len(a.ARGV) > n {
		return a.ARGV[n]
	} else {
		return ""
	}
}

func (a Args) Word(n int) string {
	for {
		arg := a.At(n)
		if arg != "" && strings.HasPrefix(arg, "-") {
			n += 1
			continue
		}
		return arg
	}
}

func (a Args) Required(n int) string {
	value := a.At(n)
	if value == "" {
		Errorln(HelpText(a.CommandName(), a.ProgramName()))
		Exit(1)
	}
	return value
}

func (a Args) HasFlag(flag string) bool {
	for _, arg := range a.ARGV {
		if arg == flag {
			return true
		}
	}
	return false
}
