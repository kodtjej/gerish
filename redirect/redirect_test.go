package redirect

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/urfave/cli"
	"gitlab.com/xonvanetta/gerish/httpserver"
)

const URL = "http://www.google.com"

func TestDefaultRedirect(t *testing.T) {
	recorder := httpserver.NewTest(cliFlags()[0], redirect)

	assert.Equal(t, URL, cliFlags()[0].(cli.StringFlag).Value)
	assert.Equal(t, http.StatusTemporaryRedirect, recorder.Result().StatusCode)

}
func TestCustomRedirect(t *testing.T) {
	customURL := "/custom-url.custom-tld"
	stringFlag := cliFlags()[0].(cli.StringFlag)
	stringFlag.Value = customURL

	recorder := httpserver.NewTest(stringFlag, redirect)

	assert.Equal(t, customURL, recorder.HeaderMap.Get("location"))
	assert.Equal(t, http.StatusTemporaryRedirect, recorder.Result().StatusCode)

}
