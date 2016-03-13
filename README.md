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
    16	func run(ctx *cli.Context, argv *argT) error {
    17		//TODO: do something here
    18		return nil
    19	}
    20	
    21	func main() {
    22		cli.Run(new(argT), func(ctx *cli.Context) error {
    23			argv := ctx.Argv().(*argT)
    24			if argv.Help {
    25				ctx.WriteUsage()
    26				return nil
    27			}
    28			return run(ctx, argv)
    29		})
    30	}
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
     6	
     7		"github.com/mkideal/cli"
     8	)
     9	
    10	func main() {
    11		if err := cli.Root(root,
    12			cli.Tree(help),
    13			cli.Tree(version),
    14		).Run(os.Args[1:]); err != nil {
    15			fmt.Fprintln(os.Stderr, err)
    16			os.Exit(1)
    17		}
    18	}
    19	
    20	//--------------
    21	// root command
    22	//--------------
    23	
    24	type rootT struct {
    25		cli.Helper
    26	}
    27	
    28	var root = &cli.Command{
    29		Name: os.Args[0],
    30		//Desc: "describe the app",
    31		Argv: func() interface{} { return new(rootT) },
    32	
    33		Fn: func(ctx *cli.Context) error {
    34			argv := ctx.Argv().(*rootT)
    35			if argv.Help || len(ctx.Args()) == 0 {
    36				ctx.WriteUsage()
    37				return nil
    38			}
    39	
    40			//TODO: do something
    41			return nil
    42		},
    43	}
    44	
    45	//--------------
    46	// help command
    47	//--------------
    48	
    49	var help = cli.HelpCommand("display help")
    50	
    51	//-----------------
    52	// version command
    53	//-----------------
    54	
    55	const appVersion = "v0.0.1"
    56	
    57	var version = &cli.Command{
    58		Name: "version",
    59		Desc: "display version",
    60	
    61		Fn: func(ctx *cli.Context) error {
    62			ctx.String(appVersion + "\n")
    63			return nil
    64		},
    65	}
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
