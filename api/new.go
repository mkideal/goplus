package api

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/mkideal/cli"
	"github.com/mkideal/goplus/etc"
)

func New() *cli.Command {
	return new_
}

type newT struct {
	cli.Helper
	Binary bool   `cli:"b,bin" usage:"generate binary application or not" dft:"true"`
	Name   string `cli:"*name" usage:"name of created project" name:"NAME"`
	Dir    string `cli:"d,dir" usage:"dir of project" dft:"./"`
	Type   string `cli:"t,type" usage:"type of project, empty or one of basic/tree/http/my" dft:"basic" name:"TYPE"`
	TplDir string `cli:"tpl-dir" usage:"useful if TYPE is empty or equal 'my'" name:"TPL_DIR"`
	Values string `cli:"values" usage:"useful if TYPE is empty. Format: --values={k1:v1/k2:v2}"`
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
	Desc: "create a application or library",
	Text: "", //TODO: add detail description

	Argv: func() interface{} { return new(newT) },

	Fn: func(ctx *cli.Context) error {
		argv := ctx.Argv().(*newT)
		if argv.Help {
			ctx.String(ctx.Usage())
			return nil
		}
		switch argv.Type {
		case etc.Basic:
			return newBasicApp(ctx, argv)

		case etc.Tree:
			return newTreeApp(ctx, argv)

		case etc.HTTP:
			return newHTTPApp(ctx, argv)

		case etc.My:
			return newMyApp(ctx, argv)
		}
		return nil
	},
}

//-------------
// type: basic
//-------------

func newBasicApp(ctx *cli.Context, argv *newT) error {
	projectDir := filepath.Join(argv.Dir, argv.Name)

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

func (argv *argT) Validate() error {
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

//------------
// type: tree
//------------

func newTreeApp(ctx *cli.Context, argv *newT) error {
	projectDir := filepath.Join(argv.Dir, argv.Name)

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
	"fmt"
	"os"
	"strings"

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
			ctx.String(ctx.Usage())
			return nil
		}

		//TODO: do something
		return nil
	},
}

//--------------
// help command
//--------------

var help = &cli.Command{
	Name:        "help",
	Desc:        "display help",
	CanSubRoute: true,

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
			return fmt.Errorf("command %s not found", clr.Yellow(strings.Join(args, " ")))
		}
		ctx.String(child.Usage(ctx))
		return nil
	},
}

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

//------------
// type: http
//------------

func newHTTPApp(ctx *cli.Context, argv *newT) error {
	projectDir := filepath.Join(argv.Dir, argv.Name)

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
	HTTPMethods: []string{http.MethodGet},

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

func (t *daemonT) Validate() error {
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

//----------
// type: my
//----------

func newMyApp(ctx *cli.Context, argv *newT) error {
	return fmt.Errorf("Not implements")
}
