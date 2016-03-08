package api

import (
	"github.com/mkideal/cli"
)

func Help() *cli.Command {
	return help
}

var help = cli.HelpCommand("display help, try `goplus help <command>' for detail")
