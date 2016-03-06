package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/mkideal/cli"
	"github.com/mkideal/goplus/api"
)

func main() {
	if err := cli.Root(api.Root(),
		cli.Tree(api.Help()),
		cli.Tree(api.Version()),
		cli.Tree(api.New()),
		cli.Tree(api.Publish()),
		bindGo("build", "compile packages and dependencies"),
		bindGo("clean", "remove object files"),
		bindGo("doc", "show documentation for package or symbol"),
		bindGo("env", "print Go environment information"),
		bindGo("fix", "run go tool fix on packages"),
		bindGo("fmt", "run gofmt on package sources"),
		bindGo("generate", "generate Go files by processing source"),
		bindGo("get", "download and install packages and dependencies"),
		bindGo("install", "compile and install packages and dependencies"),
		bindGo("list", "list packages"),
		bindGo("run", "compile and run Go program"),
		bindGo("test", "test packages"),
		bindGo("tool", "run specified go tool"),
		bindGo("vet", "run go tool vet on packages"),
	).Run(os.Args[1:]); err != nil {
		fmt.Println(err)
	}
}

func bindGo(name string, desc string) *cli.CommandTree {
	return cli.Tree(&cli.Command{
		Name:        name,
		Desc:        desc + "(go " + name + ")",
		CanSubRoute: true,

		Fn: func(ctx *cli.Context) error {
			args := append([]string{name}, ctx.Args()...)
			cmd := exec.Command("go", args...)
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			cmd.Stdin = os.Stdin
			cmd.Run()
			return nil
		},
	})
}
