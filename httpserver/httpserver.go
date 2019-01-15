package httpserver

import (
	"math"
	"net/http"
	"sync/atomic"
	"time"

	"github.com/jonaz/ginlogrus"

	"github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
	"github.com/urfave/cli"
)

var request uint64

var OKMessage = gin.H{"message": "ok"}

//CLIFunc the func that is parsed down to the httpserver
type CLIFunc func(*cli.Context, *gin.Context)

//New applies the callback to the gin instance and returns the error from gin
func New(c *cli.Context, callback CLIFunc) error {
	g := gin.New()

	g.Use(gin.Recovery())

	logger, err := getLogger(c)
	if err != nil {
		return err
	}

	g.Use(ginlogrus.New(logger, time.RFC3339))

	g.GET(c.String("url"), unstable(c, callback))
	g.POST(c.String("url"), unstable(c, callback))
	g.PUT(c.String("url"), unstable(c, callback))
	g.DELETE(c.String("url"), unstable(c, callback))

	return g.Run(":" + c.String("port"))
}

func unstable(c *cli.Context, callback CLIFunc) gin.HandlerFunc {
	return func(g *gin.Context) {
		atomic.AddUint64(&request, 1)
		rem := math.Mod(float64(atomic.LoadUint64(&request)), float64(c.Int("interval")))
		if rem == 0 {
			g.JSON(http.StatusOK, OKMessage)
			return
		}

		callback(c, g)
	}
}

func getLogger(c *cli.Context) (*logrus.Logger, error) {
	logger := logrus.StandardLogger()

	if c.Bool("json-logging") {
		logger.Formatter = &logrus.JSONFormatter{}
	}
	loglevel, err := logrus.ParseLevel(c.String("log-level"))

	if err != nil {
		return nil, err
	}

	logger.SetLevel(loglevel)

	return logger, nil
}

//CLIFlags the flags for the default httpserver
func CLIFlags(urlPath string) []cli.Flag {
	return []cli.Flag{
		cli.StringFlag{
			Name:  "port, p",
			Value: "8080",
			Usage: "The port to start the httpserver on",
		},
		cli.StringFlag{
			Name:  "url",
			Value: urlPath,
			Usage: "The url path on which it should respond on",
		},
		cli.IntFlag{
			Name:  "interval, i",
			Value: 0,
			Usage: "How often the request should return an ok message, i.e. every (n)'th request",
		},
		cli.StringFlag{
			Name:  "log-level",
			Value: "info",
			Usage: "Log level to set on the httpserver (info, warn, error)",
		},
		cli.BoolFlag{
			Name:  "json-logging",
			Usage: "If you wants the logs to be in json",
		},
	}
}
