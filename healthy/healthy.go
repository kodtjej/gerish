package healthy

import (
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
	"gitlab.com/xonvanetta/gerish/httpserver"
)

var CLICommand = cli.Command{
	Name:   "healthy",
	Usage:  "starts a healthy http server that returns an optional body",
	Action: action,
	Flags:  cliFlags(),
}

func action(c *cli.Context) error {
	return httpserver.New(c, healthy)
}

func healthy(c *cli.Context, g *gin.Context) {
	if c.String("file") == "" {
		g.String(http.StatusOK, c.String("body"))
		return
	}
	f, err := ioutil.ReadFile(c.String("file"))
	if err != nil {
		logrus.Errorf("Unable to read body from file")
		os.Exit(1)
		return
	}

	g.String(http.StatusOK, string(f))
	return
}

func cliFlags() []cli.Flag {
	return append(
		[]cli.Flag{
			cli.StringFlag{
				Name:  "body, b",
				Value: `{"message":"ok"}`,
				Usage: "Which JSON body the server should return",
			},
			cli.StringFlag{
				Name:  "file, f",
				Usage: "mycooldata.json",
			},
		},
		httpserver.CLIFlags("healthy")...,
	)
}
