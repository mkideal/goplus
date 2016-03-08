package templates

import (
	"fmt"

	"github.com/mkideal/cli"
)

var _ = register("my", My)

func My(ctx *cli.Context, cfg TemplateConfig) error {
	return fmt.Errorf("not implements")
}
