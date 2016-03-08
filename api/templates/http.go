package templates

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/mkideal/cli"
)

var _ = register("http", HTTP)

func HTTP(ctx *cli.Context, cfg TemplateConfig) error {
	projectDir := filepath.Join(cfg.Dir, cfg.Name)

	// create dir
	if err := os.Mkdir(projectDir, 0755); err != nil {
		return err
	}

	mainFile, err := os.OpenFile(filepath.Join(projectDir, "main.go"), os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}

	_, err = fmt.Fprintf(mainFile, `package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/mkideal/cli"
)

func main() {
	if err := cli.Root(root,
		cli.Tree(help),
		cli.Tree(daemon),
		cli.Tree(api,
			cli.Tree(ping),
		),
	).Run(os.Args[1:]); err != nil {
		fmt.Println(err)
	}
}

//------
// root
//------
var root = &cli.Command{
	Fn: func(ctx *cli.Context) error {
		ctx.String(ctx.Usage())
		return nil
	},
}

//------
// help
//------
var help = &cli.Command{
	Name:        "help",
	Desc:        "display help",
	CanSubRoute: true,
	HTTPRouters: []string{"/v1/help"},
	HTTPMethods: []string{"GET"},

	Fn: func(ctx *cli.Context) error {
		var (
			args   = ctx.Args()
			parent = ctx.Command().Parent()
		)
		if len(args) == 0 {
			ctx.String(parent.Usage(ctx))
			return nil
		}
		var (
			child = parent.Route(args)
			clr   = ctx.Color()
		)
		if child == nil {
			return fmt.Errorf("command %%s not found", clr.Yellow(strings.Join(args, " ")))
		}
		ctx.String(child.Usage(ctx))
		return nil
	},
}

//--------
// daemon
//--------
type daemonT struct {
	cli.Helper
	Port uint16 %scli:"p,port" usage:"http port" dft:"8080"%s
}

func (t *daemonT) Validate(ctx *cli.Context) error {
	if t.Port == 0 {
		return fmt.Errorf("please don't use 0 as http port")
	}
	return nil
}

var daemon = &cli.Command{
	Name: "daemon",
	Desc: "startup app as daemon",
	Argv: func() interface{} { return new(daemonT) },
	Fn: func(ctx *cli.Context) error {
		argv := ctx.Argv().(*daemonT)
		if argv.Help {
			ctx.String(ctx.Usage())
			return nil
		}

		//NOTE: remove following line will disable debug mode
		cli.EnableDebug()

		addr := fmt.Sprintf(":%%d", argv.Port)
		cli.Debugf("http addr: %%s", addr)

		if err := ctx.Command().Root().RegisterHTTP(ctx); err != nil {
			return err
		}
		return http.ListenAndServe(addr, ctx.Command().Root())
	},
}

//-----
// api
//-----
var api = &cli.Command{
	Name: "api",
	Desc: "display all api",
	Fn: func(ctx *cli.Context) error {
		ctx.String("Commands:\n")
		ctx.String("    ping\n")
		return nil
	},
}

//------
// ping
//------
var ping = &cli.Command{
	Name: "ping",
	Desc: "ping server",
	Fn: func(ctx *cli.Context) error {
		ctx.String("pong\n")
		return nil
	},
}`, "`", "`")
	return err
}
