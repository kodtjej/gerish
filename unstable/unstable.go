package unstable

import (
	"math"
	"net/http"
	"sync/atomic"

	"github.com/gin-gonic/gin"
	"github.com/urfave/cli"
	"gitlab.com/xonvanetta/gerish/httpserver"
)

var count int64

//CLI CLI
func CLI(c *cli.Context) error {
	return httpserver.New(c, unstable)
}

func unstable(c *cli.Context) gin.HandlerFunc {
	return func(g *gin.Context) {
		atomic.AddInt64(&count, 1)

		if math.Mod(float64(atomic.LoadInt64(&count)), float64(c.Int("interval"))) == 0 {
			g.AbortWithStatus(http.StatusRequestTimeout)
			return
		}

		g.JSON(http.StatusOK, gin.H{"Message": "Ok"})
	}
}

//CLIFlags the flags for the webserver
func CLIFlags() []cli.Flag {
	return append(
		[]cli.Flag{
			cli.IntFlag{
				Name:  "interval, i",
				Value: 2,
				Usage: "How often the request should return an error",
			},
		},
		httpserver.CLIFlags("unstable")...,
	)
}
