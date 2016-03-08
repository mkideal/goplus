package api

import (
	"fmt"
	"strings"

	"github.com/mkideal/cli"
	"github.com/mkideal/goplus/api/templates"
)

func New() *cli.Command {
	return new_
}

type newT struct {
	cli.Helper
	Type string `cli:"t,type" usage:"type of project" dft:"basic" name:"TYPE"`
	List bool   `cli:"!l,list" usage:"list all types of project template"`
	templates.TemplateConfig
}

func (t *newT) Validate(ctx *cli.Context) error {
	clr := ctx.Color()
	b := clr.Bold
	if t.Name == "" {
		return fmt.Errorf("%s is empty", b("NAME"))
	}
	if t.Type == "" && t.TplDir == "" {
		return fmt.Errorf("%s and %s both are empty", b("TYPE"), b("TPL_DIR"))
	}
	if t.Type != "" && !templates.ValidateType(t.Type) {
		return fmt.Errorf("%s is invalid, try `goplus new -l'", b("TYPE"), b(t.Type))
	}
	return nil
}

var new_ = &cli.Command{
	Name: "new",
	Desc: "create application skeleton by template type",
	Argv: func() interface{} { return new(newT) },

	Fn: func(ctx *cli.Context) error {
		argv := ctx.Argv().(*newT)
		if argv.Help {
			ctx.String(ctx.Usage())
			return nil
		}
		if argv.List {
			prefix := "\t"
			content := strings.Join(templates.List(), "\n"+prefix)
			ctx.String("list of all types:\n" + prefix + content + "\n")
			return nil
		}
		return templates.New(argv.Type, ctx, argv.TemplateConfig)
	},
}
