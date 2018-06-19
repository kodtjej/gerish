package httpserver

import (
	"math"
	"net/http"
	"sync/atomic"

	"github.com/gin-gonic/gin"
	"github.com/urfave/cli"
)

var count int64

//CLIFunc the func that is parsed down to the httpserver
type CLIFunc func(*cli.Context, *gin.Context)

//New applies the callback to the gin instance and returns the error from gin
func New(c *cli.Context, callback CLIFunc) error {
	g := gin.Default()

	g.GET(c.String("url"), unstable(c, callback))
	g.POST(c.String("url"), unstable(c, callback))
	g.PUT(c.String("url"), unstable(c, callback))
	g.DELETE(c.String("url"), unstable(c, callback))

	return g.Run(":" + c.String("port"))
}

func unstable(c *cli.Context, callback CLIFunc) gin.HandlerFunc {
	return func(g *gin.Context) {
		atomic.AddInt64(&count, 1)

		if math.Mod(float64(atomic.LoadInt64(&count)), float64(c.Int("interval"))) == 0 {
			g.JSON(http.StatusOK, gin.H{"Message": "Ok"})
			return
		}

		callback(c, g)
	}
}

//CLIFlags the flags for the default webserver
func CLIFlags(urlPath string) []cli.Flag {
	return []cli.Flag{
		cli.StringFlag{
			Name:  "port, p",
			Value: "8080",
			Usage: "The port to start the webserver on",
		},
		cli.StringFlag{
			Name:  "url",
			Value: urlPath,
			Usage: "The url path on which it should respond on",
		},
		cli.IntFlag{
			Name:  "interval, i",
			Value: 1,
			Usage: "How often the request should return an ok message",
		},
	}
}
