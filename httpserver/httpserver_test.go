package httpserver

import (
	"flag"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/urfave/cli"
)

func TestHttpServer(t *testing.T) {
	intervalFlag := CLIFlags("interval")[2].(cli.IntFlag)
	intervalFlag.Value = 3

	tests := []struct {
		code    int
		message string
	}{
		{http.StatusNotFound, "{\"Message\":\"Not Found\"}"},
		{http.StatusNotFound, "{\"Message\":\"Not Found\"}"},
		{http.StatusOK, DefaultOkString},
		{http.StatusNotFound, "{\"Message\":\"Not Found\"}"},
		{http.StatusNotFound, "{\"Message\":\"Not Found\"}"},
		{http.StatusOK, DefaultOkString},
		{http.StatusNotFound, "{\"Message\":\"Not Found\"}"},
	}

	app := cli.NewApp()

	flags := flag.NewFlagSet("test", 0)
	//Apply the cliFlags to the context of cli
	intervalFlag.Apply(flags)

	cliContext := cli.NewContext(app, flags, nil)

	g := gin.New()
	req, _ := http.NewRequest("GET", "/test", nil)

	g.GET("/test", unstable(cliContext, func(c *cli.Context, g *gin.Context) {
		g.JSON(http.StatusNotFound, gin.H{"Message": "Not Found"})
	}))

	for _, test := range tests {
		recorder := httptest.NewRecorder()
		g.ServeHTTP(recorder, req)
		assert.Equal(t, test.code, recorder.Result().StatusCode)
		assert.Equal(t, test.message, recorder.Body.String())
	}
}

func TestCLIFlags(t *testing.T) {
	url := "testFlags"
	cliURLFLag := CLIFlags(url)[1].(cli.StringFlag)

	assert.Equal(t, cliURLFLag.Value, url)
}
