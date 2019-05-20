package main

import (
	"os"

	"github.com/urfave/cli"
	"gitlab.com/xonvanetta/gerish/connectionfail"
	"gitlab.com/xonvanetta/gerish/faulty"
	"gitlab.com/xonvanetta/gerish/healthy"
	"gitlab.com/xonvanetta/gerish/redirect"
	"gitlab.com/xonvanetta/gerish/slow"
)

func main() {
	app := cli.NewApp()
	app.Commands = []cli.Command{
		slow.CLICommand,
		faulty.CLICommand,
		redirect.CLICommand,
		connectionfail.CLICommand,
		healthy.CLICommand,
	}

	err := app.Run(os.Args)
	if err != nil {
		panic(err)
	}
}
