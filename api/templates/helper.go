package templates

import (
	"fmt"
	"sort"

	"github.com/mkideal/cli"
)

type TemplateConfig struct {
	Name   string `cli:"*name" usage:"name of created project" name:"NAME"`
	Dir    string `cli:"d,dir" usage:"parent dir of generated project" dft:"./"`
	TplDir string `cli:"tpl-dir" usage:"useful if TYPE is empty or equal to my" name:"TPL_DIR"`
	Values string `cli:"values" usage:"useful if TYPE is empty. format: --values={k1:v1/k2:v2}"`
}

type maker func(*cli.Context, TemplateConfig) error

var templatesMap = make(map[string]maker)

func register(typ string, fn maker) bool {
	if _, ok := templatesMap[typ]; ok {
		cli.Panicf("repeat register template %s", typ)
	}
	templatesMap[typ] = fn
	return true
}

func New(typ string, ctx *cli.Context, cfg TemplateConfig) error {
	clr := ctx.Color()
	fn, ok := templatesMap[typ]
	if !ok {
		return fmt.Errorf("unsupported template type %s", clr.Yellow(typ))
	}
	return fn(ctx, cfg)
}

func ValidateType(typ string) bool {
	_, ok := templatesMap[typ]
	return ok
}

func List() []string {
	list := make([]string, 0, len(templatesMap))
	for t, _ := range templatesMap {
		list = append(list, t)
	}
	sort.Strings(list)
	return list
}
