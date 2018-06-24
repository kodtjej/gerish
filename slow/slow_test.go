package slow

import (
	"net/http"
	"testing"
	"time"

	"gitlab.com/xonvanetta/gerish/httpserver"

	"github.com/stretchr/testify/assert"
	"github.com/urfave/cli"
)

func TestSlow(t *testing.T) {
	start := time.Now()
	recorder := httpserver.NewTest(cliFlags()[0], slow)

	end := time.Since(start)

	assert.Equal(t, cliFlags()[0].(cli.IntFlag).Value, int(end.Seconds()))
	assert.Equal(t, httpserver.DefaultOkString, recorder.Body.String())
	assert.Equal(t, http.StatusOK, recorder.Result().StatusCode)
}

func TestSlowWithLessSeconds(t *testing.T) {
	seconds := 1
	intFlag := cliFlags()[0].(cli.IntFlag)

	intFlag.Value = seconds

	start := time.Now()
	recorder := httpserver.NewTest(intFlag, slow)
	end := time.Since(start)

	assert.Equal(t, seconds, int(end.Seconds()))
	assert.Equal(t, httpserver.DefaultOkString, recorder.Body.String())
	assert.Equal(t, http.StatusOK, recorder.Result().StatusCode)
}
