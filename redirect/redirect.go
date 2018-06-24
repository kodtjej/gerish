package redirect

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/urfave/cli"
	"gitlab.com/xonvanetta/gerish/httpserver"
)

//CLICommand CLICommand
var CLICommand = cli.Command{
	Name:    "redirect",
	Aliases: []string{"r"},
	Usage:   "starts a http server redirects to different URL",
	Action:  action,
	Flags:   cliFlags(),
}

func action(c *cli.Context) error {
	return httpserver.New(c, redirect)
}

func redirect(c *cli.Context, g *gin.Context) {
	g.Redirect(http.StatusTemporaryRedirect, c.String("redirect-url"))
}

func cliFlags() []cli.Flag {
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
