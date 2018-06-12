package httpserver

import (
	"github.com/gin-gonic/gin"
	"github.com/urfave/cli"
)

//CLIFlags the flags for the default webserver
var CLIFlags = []cli.Flag{
	cli.StringFlag{
		Name:  "port, p",
		Value: "8080",
		Usage: "The port to start the webserver on",
	},
}

//New returns a new gin instance
func New() *gin.Engine {
	return gin.Default()
}

//Run runs the webserver
func Run(g *gin.Engine, c *cli.Context) error {
	return g.Run(":" + c.String("port"))
}
