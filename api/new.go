package api

import (
	"fmt"

	"github.com/mkideal/cli"
	"github.com/mkideal/goplus/api/templates"
	"github.com/mkideal/goplus/etc"
)

func New() *cli.Command {
	return new_
}

type newT struct {
	cli.Helper
	Type string `cli:"t,type" usage:"type of project, empty or one of basic/tree/http/my" dft:"basic" name:"TYPE"`
	templates.TemplateConfig
}

func (t *newT) Validate() error {
	if t.Name == "" {
		return fmt.Errorf("NAME is empty")
	}
	if t.Type == "" && t.TplDir == "" {
		return fmt.Errorf("TYPE and TPL_DIR both are empty")
	}
	if t.Type != "" && !etc.ValidateType(t.Type) {
		return fmt.Errorf("TYPE invalid, try `goplus new -h'")
	}
	return nil
}

var new_ = &cli.Command{
	Name: "new",
	Desc: "create a application",
	Text: "", //TODO: add detail description

	Argv: func() interface{} { return new(newT) },

	Fn: func(ctx *cli.Context) error {
		argv := ctx.Argv().(*newT)
		if argv.Help {
			ctx.String(ctx.Usage())
			return nil
		}
		err, ok := templates.New(ctx, argv.Type, argv.TemplateConfig)
		if !ok {
			return newMyApp(ctx, argv.TemplateConfig)
		}
		return err
	},
}

//----------
// type: my
//----------

func newMyApp(ctx *cli.Context, argv templates.TemplateConfig) error {
	return fmt.Errorf("Not implements")
}
