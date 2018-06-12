package httpserver

import (
	"github.com/gin-gonic/gin"
	"github.com/urfave/cli"
)

//CLIFunc the func that is parsed down to the httpserver
type CLIFunc func(*cli.Context) gin.HandlerFunc

//New applies the callback to the gin instance and returns the error from gin
func New(c *cli.Context, callback CLIFunc) error {
	g := gin.Default()

	g.GET(c.String("url"), callback(c))
	g.POST(c.String("url"), callback(c))
	g.PUT(c.String("url"), callback(c))
	g.DELETE(c.String("url"), callback(c))

	return g.Run(":" + c.String("port"))
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
	}
}
