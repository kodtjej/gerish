package faulty

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/urfave/cli"
	"gitlab.com/xonvanetta/gerish/httpserver"
)

//CLI CLI
func CLI(c *cli.Context) error {
	return httpserver.New(c, faulty)
}

func faulty(c *cli.Context, g *gin.Context) {
	g.AbortWithStatus(c.Int("code"))
}

//CLIFlags the flags for the webserver
func CLIFlags() []cli.Flag {
	return append(
		[]cli.Flag{
			cli.IntFlag{
				Name:  "code, c",
				Value: http.StatusRequestTimeout,
				Usage: "Which HTTP status code the server should return on error",
			},
		},
		httpserver.CLIFlags("faulty")...,
	)
}
