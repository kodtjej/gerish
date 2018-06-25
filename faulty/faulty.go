package faulty

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/urfave/cli"
	"gitlab.com/xonvanetta/gerish/httpserver"
)

//CLICommand CLICommand
var CLICommand = cli.Command{
	Name:    "faulty",
	Aliases: []string{"f"},
	Usage:   "starts a faulty http server that returns error codes",
	Action:  action,
	Flags:   cliFlags(),
}

func action(c *cli.Context) error {
	return httpserver.New(c, faulty)
}

func faulty(c *cli.Context, g *gin.Context) {
	g.AbortWithStatus(c.Int("code"))
}

func cliFlags() []cli.Flag {
	return append(
		[]cli.Flag{
			cli.IntFlag{
				Name:  "code, c",
				Value: http.StatusNotFound,
				Usage: "Which HTTP status code the server should return on error",
			},
		},
		httpserver.CLIFlags("faulty")...,
	)
}
