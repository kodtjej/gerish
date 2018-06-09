package unstable

import (
	"fmt"

	"github.com/urfave/cli"
)

//CLI CLI
func CLI(c *cli.Context) error {
	fmt.Println("starts a http server that is unstable and might return error codes")
	return nil
}
