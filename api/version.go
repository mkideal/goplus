package api

import (
	"os/exec"
	"strings"

	"github.com/mkideal/cli"
	"github.com/mkideal/goplus/etc"
)

func Version() *cli.Command {
	return version
}

var version = &cli.Command{
	Name: "version",
	Desc: "display goplus version",

	Fn: func(ctx *cli.Context) error {
		ctx.String("goplus version: " + etc.GoplusVersion + "\n")
		cmd := exec.Command("go", "version")
		reader, err := cmd.StdoutPipe()
		if err != nil {
			return err
		}
		defer reader.Close()
		cmd.Start()
		buf := make([]byte, 256)
		if n, err := reader.Read(buf); err != nil {
			return err
		} else {
			v := strings.TrimPrefix(string(buf[:n]), "go version")
			v = strings.TrimSuffix(v, "\n")
			ctx.String("golang version:" + v + "\n")
		}
		cmd.Wait()
		return nil
	},
}
