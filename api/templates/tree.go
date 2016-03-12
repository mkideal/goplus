package templates

import (
	"os"
	"path/filepath"

	"github.com/mkideal/cli"
)

var _ = register("tree", Tree)

func Tree(ctx *cli.Context, cfg TemplateConfig) error {
	projectDir := filepath.Join(cfg.Dir, cfg.Name)

	// create dir
	if err := os.Mkdir(projectDir, 0755); err != nil {
		return err
	}

	filename := filepath.Join(projectDir, "main.go")
	mainFile, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}

	_, err = mainFile.WriteString(`package main

import (
	"fmt"
	"os"

	"github.com/mkideal/cli"
)

func main() {
	if err := cli.Root(root,
		cli.Tree(help),
		cli.Tree(version),
	).Run(os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

//--------------
// root command
//--------------

type rootT struct {
	cli.Helper
}

var root = &cli.Command{
	Name: os.Args[0],
	//Desc: "describe the app",
	Argv: func() interface{} { return new(rootT) },

	Fn: func(ctx *cli.Context) error {
		argv := ctx.Argv().(*rootT)
		if argv.Help || len(ctx.Args()) == 0 {
			ctx.WriteUsage()
			return nil
		}

		//TODO: do something
		return nil
	},
}

//--------------
// help command
//--------------

var help = cli.HelpCommand("display help")

//-----------------
// version command
//-----------------

const appVersion = "v0.0.1"

var version = &cli.Command{
	Name: "version",
	Desc: "display version",

	Fn: func(ctx *cli.Context) error {
		ctx.String(appVersion + "\n")
		return nil
	},
}`)
	return err
}
