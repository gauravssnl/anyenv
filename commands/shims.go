package commands

import (
	"github.com/mislav/anyenv/cli"
	"github.com/mislav/anyenv/config"
)

var shimsHelp = `
Usage: $ProgramName shims [--short]

List existing $ProgramName shims
`

func shimsCmd(args cli.Args) {
	shimsDir := config.ShimsDir()
	short := args.HasFlag("--short")

	for _, shim := range shimsDir.Entries() {
		if short {
			cli.Println(shim.Base())
		} else {
			cli.Println(shim)
		}
	}
}

func init() {
	cli.Register("shims", shimsCmd, shimsHelp)
}
