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
	g := gin.Default()

	g.GET(c.String("url"), slow(c.Int("second")))
	g.POST(c.String("url"), slow(c.Int("second")))
	g.PUT(c.String("url"), slow(c.Int("second")))
	g.DELETE(c.String("url"), slow(c.Int("second")))

	return g.Run(":" + c.String("port"))
}

func slow(seconds int) gin.HandlerFunc {
	return func(g *gin.Context) {
		time.Sleep(time.Second * time.Duration(seconds))
		g.JSON(http.StatusOK, gin.H{"Message": "Ok"})
	}
}

//CLIFlags the flags for the webserver
func CLIFlags() []cli.Flag {
	return append(
		[]cli.Flag{
			cli.IntFlag{
				Name:  "second, s",
				Value: 11,
				Usage: "How long should the request take before continuing",
			},
			cli.StringFlag{
				Name:  "url",
				Value: "/slow",
				Usage: "The url path on which it should respond on",
			},
		},
		httpserver.CLIFlags...,
	)
}
