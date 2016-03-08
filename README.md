# goplus [![License](http://img.shields.io/badge/license-mit-blue.svg?style=flat-square)](https://raw.githubusercontent.com/mkideal/goplus/master/LICENSE)

## License

[The MIT License (MIT)](https://raw.githubusercontent.com/mkideal/goplus/master/LICENSE)

## Install
```sh
go get github.com/mkideal/goplus
```

## TODOs

* More application template
* Complete `publish` command

## Getting started

`goplus` binds all commands of local go program. So you can use goplus <command> instead of go <command>, e.g.

	goplus build
	goplus install
	goplus test
	goplus run
	goplus fmt
	......

`goplus` has some new commands, e.g.

	goplus new
	goplus publish
	......

### new command

`new` command creates a protect. Following command will generate a dir `hello` at current dir. Dir `hello` contains one file: `main.go`.

```shell
$> goplus new --name hello
```

```shell
$> cd hello
$> cat -n main.go
     1	package main
     2	
     3	import (
     4		"github.com/mkideal/cli"
     5	)
     6	
     7	type argT struct {
     8		cli.Helper
     9	}
    10	
    11	func (argv *argT) Validate(ctx *cli.Context) error {
    12		//TODO: validate something or remove this function.
    13		return nil
    14	}
    15	
    16	func main() {
    17		cli.Run(new(argT), func(ctx *cli.Context) error {
    18			argv := ctx.Argv().(*argT)
    19			if argv.Help {
    20				ctx.String(ctx.Usage())
    21				return nil
    22			}
    23	
    24			//TODO: remove following line, and do something here
    25			ctx.JSONIndentln(argv, "", "    ")
    26	
    27			return nil
    28		})
    29	}
```

You can specify what type your application. The `TYPE` must be one of these:

	basic
	tree
	http

default is `basic`.

Following command will generate a command tree application.

```shell
$> goplus new -t tree --name treesample
```

```shell
$> cd treesample
$> cat -n main.go
     1	package main
     2	
     3	import (
     4		"fmt"
     5		"os"
     6		"strings"
     7	
     8		"github.com/mkideal/cli"
     9	)
    10	
    11	func main() {
    12		if err := cli.Root(root,
    13			cli.Tree(help),
    14			cli.Tree(version),
    15		).Run(os.Args[1:]); err != nil {
    16			fmt.Fprintln(os.Stderr, err)
    17			os.Exit(1)
    18		}
    19	}
    20	
    21	//--------------
    22	// root command
    23	//--------------
    24	
    25	type rootT struct {
    26		cli.Helper
    27	}
    28	
    29	var root = &cli.Command{
    30		Name: os.Args[0],
    31		//Desc: "describe the app",
    32		Argv: func() interface{} { return new(rootT) },
    33	
    34		Fn: func(ctx *cli.Context) error {
    35			argv := ctx.Argv().(*rootT)
    36			if argv.Help || len(ctx.Args()) == 0 {
    37				ctx.String(ctx.Usage())
    38				return nil
    39			}
    40	
    41			//TODO: do something
    42			return nil
    43		},
    44	}
    45	
    46	//--------------
    47	// help command
    48	//--------------
    49	
    50	var help = &cli.Command{
    51		Name:        "help",
    52		Desc:        "display help",
    53		CanSubRoute: true,
    54	
    55		Fn: func(ctx *cli.Context) error {
    56			var (
    57				args   = ctx.Args()
    58				parent = ctx.Command().Parent()
    59			)
    60			if len(args) == 0 {
    61				ctx.String(parent.Usage(ctx))
    62				return nil
    63			}
    64			var (
    65				child = parent.Route(args)
    66				clr   = ctx.Color()
    67			)
    68			if child == nil {
    69				return fmt.Errorf("command %s not found", clr.Yellow(strings.Join(args, " ")))
    70			}
    71			ctx.String(child.Usage(ctx))
    72			return nil
    73		},
    74	}
    75	
    76	//-----------------
    77	// version command
    78	//-----------------
    79	
    80	const appVersion = "v0.0.1"
    81	
    82	var version = &cli.Command{
    83		Name: "version",
    84		Desc: "display version",
    85	
    86		Fn: func(ctx *cli.Context) error {
    87			ctx.String(appVersion + "\n")
    88			return nil
    89		},
    90	}
```

Try

```shell
$> goplus new -t http --name httpserver
$> cd httpserver
$> go build
$> ./httpserver daemon
```

```shell
$> curl http://127.0.0.1:8080/help
Commands:
  help     display help
  daemon   startup app as daemon
  api      display all api
$> curl http://127.0.0.1:8080/api
Commands:
    ping
$> curl http://127.0.0.1:8080/api/ping
pong
```
