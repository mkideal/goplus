package api

import (
	"github.com/mkideal/cli"
)

func Publish() *cli.Command {
	return publish
}

type publishT struct {
	cli.Helper
	ConfigFile string `cli:"config" usage:"config file for publish" dft:"goplus.yaml"`
}

var publish = &cli.Command{
	Name: "publish",
	Desc: "publish application or library to remote",

	Argv: func() interface{} { return new(publishT) },

	Fn: func(ctx *cli.Context) error {
		argv := ctx.Argv().(*publishT)
		if argv.Help {
			ctx.String(ctx.Usage())
			return nil
		}
		return nil
	},
}
