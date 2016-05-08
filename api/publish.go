package api

import (
	"fmt"
	"io/ioutil"

	"github.com/mkideal/cli"
)

func Publish() *cli.Command {
	return publish
}

type publishT struct {
	cli.Helper
	ConfigFile string `cli:"config" usage:"config file for publish" dft:"goplus.yaml"`
	Yes        bool   `cli:"y,yes" usage:"yes for all question" dft:"false"`

	config publishConfig `cli:"-"`
}

func (argv *publishT) Validate(ctx *cli.Context) error {
	if argv.ConfigFile == "" {
		return fmt.Errorf("config file is empty")
	}
	data, err := ioutil.ReadFile(argv.ConfigFile)
	if err != nil {
		return err
	}
	_ = data

	//TODO: unmarshal yaml to argv.config
	return nil
}

var publish = &cli.Command{
	Name: "publish",
	Desc: "publish application or package to remote",
	Argv: func() interface{} { return new(publishT) },

	Fn: func(ctx *cli.Context) error {
		//argv := ctx.Argv().(*publishT)
		return nil
	},
}

type publishConfig struct {
	Version   string
	Build     string
	Dir       string
	FileGlobs []string
}
