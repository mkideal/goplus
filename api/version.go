package api

import (
	"github.com/mkideal/cli"
	"github.com/mkideal/goplus/etc"
)

func Version() *cli.Command {
	return version
}

var version = &cli.Command{
	Name: "version",
	Desc: "display goplus version",

	Fn: func(ctx *cli.Context) error {
		ctx.String(etc.GoplusVersion + "\n")
		return nil
	},
}
