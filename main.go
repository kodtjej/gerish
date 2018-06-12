package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
	"gitlab.com/xonvanetta/gerish/faulty"
	"gitlab.com/xonvanetta/gerish/slow"
	"gitlab.com/xonvanetta/gerish/unstable"
)

func main() {
	app := cli.NewApp()
	app.Commands = []cli.Command{
		{
			Name:    "slow",
			Aliases: []string{"s"},
			Usage:   "starts a slow http server",
			Action:  slow.CLI,
			Flags:   slow.CLIFlags(),
		},
		{
			Name:    "faulty",
			Aliases: []string{"f"},
			Usage:   "starts a faulty http server that returns error codes",
			Action:  faulty.CLI,
		},
		{
			Name:    "unstable",
			Aliases: []string{"u"},
			Usage:   "starts a http server that is unstable and might return error codes",
			Action:  unstable.CLI,
		},
	}

	fmt.Println(app.Run(os.Args))
}
