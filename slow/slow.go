package slow

import (
	"net/http"
	"time"

	"gitlab.com/xonvanetta/gerish/httpserver"

	"github.com/gin-gonic/gin"
	"github.com/urfave/cli"
)

//CLICommand the command that runs the slow webserver
var CLICommand = cli.Command{
	Name:    "slow",
	Aliases: []string{"s"},
	Usage:   "starts a slow http server",
	Action:  action,
	Flags:   cliFlags(),
}

func action(c *cli.Context) error {
	return httpserver.New(c, slow)
}

func slow(c *cli.Context, g *gin.Context) {
	time.Sleep(time.Second * time.Duration(c.Int("seconds")))
	g.JSON(http.StatusOK, gin.H{"Message": "Ok"})
}

func cliFlags() []cli.Flag {
	return append(
		[]cli.Flag{
			cli.IntFlag{
				Name:  "seconds, s",
				Value: 11,
				Usage: "How long should the request take before continuing",
			},
		},
		httpserver.CLIFlags("slow")...,
	)
}
