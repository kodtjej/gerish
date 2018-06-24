package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
	"gitlab.com/xonvanetta/gerish/faulty"
	"gitlab.com/xonvanetta/gerish/redirect"
	"gitlab.com/xonvanetta/gerish/slow"
)

func main() {
	app := cli.NewApp()
	app.Commands = []cli.Command{
		slow.CLICommand,
		faulty.CLICommand,
		redirect.CLICommand,
	}

	fmt.Println(app.Run(os.Args))
}
