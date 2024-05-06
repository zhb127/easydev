package main

import (
	"os"

	"github.com/pkg/errors"
	"github.com/urfave/cli/v2"

	"github.com/zhb127/easydev/cmd/prompt"
	"github.com/zhb127/easydev/cmd/render"
	"github.com/zhb127/easydev/errs"
	"github.com/zhb127/easydev/pkg/log"
)

func main() {
	cliApp := &cli.App{
		Commands: []*cli.Command{
			render.Cmd,
			prompt.Cmd,
		},
	}

	if err := cliApp.Run(os.Args); err != nil {
		if errors.Is(err, errs.ErrInterrupt) {
			os.Exit(130)
		}
		log.Err(err)
		os.Exit(1)
	}
}
