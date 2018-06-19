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
			Flags:   faulty.CLIFlags(),
		},
		{
			Name:    "redirect",
			Aliases: []string{"r"},
			Usage:   "starts a http server redirects to different URL",
			Action:  redirect.CLI,
			Flags:   redirect.CLIFlags(),
		},
	}

	fmt.Println(app.Run(os.Args))
}
