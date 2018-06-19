package redirect

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/urfave/cli"
	"gitlab.com/xonvanetta/gerish/httpserver"
)

//CLI CLI
func CLI(c *cli.Context) error {
	return httpserver.New(c, redirect)
}

func redirect(c *cli.Context, g *gin.Context) {
	g.Redirect(http.StatusTemporaryRedirect, c.String("redirect-url"))
}

//CLIFlags the flags for the webserver
func CLIFlags() []cli.Flag {
	return append(
		[]cli.Flag{
			cli.StringFlag{
				Name:  "redirect-url, ru",
				Value: "http://www.google.com",
				Usage: "URL to redirect to",
			},
		},
		httpserver.CLIFlags("redirect")...,
	)
}
