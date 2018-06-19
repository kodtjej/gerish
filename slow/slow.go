package slow

import (
	"net/http"
	"time"

	"gitlab.com/xonvanetta/gerish/httpserver"

	"github.com/gin-gonic/gin"
	"github.com/urfave/cli"
)

//CLI CLI
func CLI(c *cli.Context) error {
	return httpserver.New(c, slow)
}

func slow(c *cli.Context, g *gin.Context) {
	time.Sleep(time.Second * time.Duration(c.Int("seconds")))
	g.JSON(http.StatusOK, gin.H{"Message": "Ok"})
}

//CLIFlags the flags for the webserver
func CLIFlags() []cli.Flag {
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
