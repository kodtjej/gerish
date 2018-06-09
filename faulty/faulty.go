package faulty

import (
	"fmt"

	"github.com/urfave/cli"
)

//CLI CLI
func CLI(c *cli.Context) error {
	fmt.Println("starts a http server that returns error codes")
	return nil
}
