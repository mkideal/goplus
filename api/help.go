package api

import (
	"fmt"
	"strings"

	"github.com/mkideal/cli"
)

func Help() *cli.Command {
	return help
}

var help = &cli.Command{
	Name:        "help",
	Desc:        "display help, try `goplus help <command>' for detail",
	CanSubRoute: true,

	Fn: func(ctx *cli.Context) error {
		args := ctx.Args()
		parent := ctx.Command().Parent()
		if len(args) == 0 {
			ctx.String(parent.Usage(ctx))
			return nil
		}
		child := parent.Route(args)
		clr := ctx.Color()
		if child == nil {
			return fmt.Errorf("command %s not found", clr.Yellow(strings.Join(args, " ")))
		}
		ctx.String(child.Usage(ctx))
		return nil
	},
}
