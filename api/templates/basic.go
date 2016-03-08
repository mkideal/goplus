package templates

import (
	"os"
	"path/filepath"

	"github.com/mkideal/cli"
)

var _ = register("basic", Basic)

func Basic(ctx *cli.Context, cfg TemplateConfig) error {
	projectDir := filepath.Join(cfg.Dir, cfg.Name)

	// create dir
	if err := os.Mkdir(projectDir, 0755); err != nil {
		return err
	}

	mainFile, err := os.OpenFile(filepath.Join(projectDir, "main.go"), os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}

	_, err = mainFile.WriteString(`package main

import (
	"github.com/mkideal/cli"
)

type argT struct {
	cli.Helper
}

func (argv *argT) Validate(ctx *cli.Context) error {
	//TODO: validate something or remove this function.
	return nil
}

func main() {
	cli.Run(new(argT), func(ctx *cli.Context) error {
		argv := ctx.Argv().(*argT)
		if argv.Help {
			ctx.String(ctx.Usage())
			return nil
		}

		//TODO: remove following line, and do something here
		ctx.JSONIndentln(argv, "", "    ")

		return nil
	})
}
`)
	return err
}
