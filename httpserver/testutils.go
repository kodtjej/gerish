package httpserver

import (
	"flag"
	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
	"github.com/urfave/cli"
)

const (
	DefaultOkString = "{\"message\":\"ok\"}"
)

//NewTest use this to create a new test case with the httpserver and the provided callback into the function.
//This should only ever be used in tests
func NewTest(cliFlag cli.Flag, callback CLIFunc) *httptest.ResponseRecorder {
	recorder := httptest.NewRecorder()
	app := cli.NewApp()

	flags := flag.NewFlagSet("test", 0)
	//Apply the cliFlags to the context of cli
	cliFlag.Apply(flags)

	cliContext := cli.NewContext(app, flags, nil)

	g := gin.New()
	req, _ := http.NewRequest("GET", "/test", nil)

	g.GET("/test", wrapper(cliContext, callback))

	g.ServeHTTP(recorder, req)

	return recorder
}

func wrapper(c *cli.Context, callback CLIFunc) gin.HandlerFunc {
	return func(g *gin.Context) {
		callback(c, g)
	}
}
