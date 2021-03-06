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
	"os"

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
		ctx.WriteUsage()
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

	Fn: cli.HelpCommandFn,
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

		addr := fmt.Sprintf(":%%d", argv.Port)
		r := ctx.Command().Root()
		if err := r.RegisterHTTP(ctx); err != nil {
			return err
		}
		return r.ListenAndServeHTTP(addr)
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
