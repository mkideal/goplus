package api

import (
	"github.com/mkideal/cli"
	"github.com/mkideal/goplus/etc"
)

func Root() *cli.Command {
	return root
}

type rootT struct {
	cli.Helper
	Version bool `cli:"v,version" usage:"display goplus version"`
}

var root = &cli.Command{
	Name: "goplus",
	Desc: "goplus is a tool for generate, build and publish golang application and library",

	Text: `goplus binds all commands of go, e.g. build,install,run,test,get,fmt...
So, you can use goplus <command> instead of go <command>, e.g.

	goplus build
	goplus install
	goplus test
	goplus run
	goplus fmt
	......

goplus has some new commands, e.g.

	goplus new
	goplus publish
	......`,

	Argv: func() interface{} { return new(rootT) },

	Fn: func(ctx *cli.Context) error {
		argv := ctx.Argv().(*rootT)
		if argv.Help || len(ctx.Args()) == 0 {
			ctx.WriteUsage()
		} else if argv.Version {
			ctx.String(etc.GoplusVersion + "\n")
		}
		return nil
	},
}
