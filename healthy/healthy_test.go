package healthy

import (
	"bytes"
	"net/http"
	"testing"

	"gitlab.com/xonvanetta/gerish/httpserver"

	"github.com/stretchr/testify/assert"
	"github.com/urfave/cli"
)

func TestHealty(t *testing.T) {
	recorder := httpserver.NewTest(cliFlags()[0], healthy)

	buffer := new(bytes.Buffer)
	buffer.ReadFrom(recorder.Result().Body)
	assert.Equal(t, http.StatusOK, recorder.Result().StatusCode)
	assert.Equal(t, "{\"message\":\"ok\"}", buffer.String())
}

func TestHealthyBodyFromFile(t *testing.T) {
	fileFlag := cliFlags()[1].(cli.StringFlag)
	fileFlag.Value = "test.json"

	recorder := httpserver.NewTest(fileFlag, healthy)

	buffer := new(bytes.Buffer)
	buffer.ReadFrom(recorder.Result().Body)
	assert.Equal(t, http.StatusOK, recorder.Result().StatusCode)
	assert.Equal(t, "{\"message\":\"fromfile\"}\n", buffer.String())
}
