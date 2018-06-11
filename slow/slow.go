package slow

import (
	"fmt"

	"github.com/urfave/cli"
)

//CLI CLI
func CLI(c *cli.Context) error {
	fmt.Println("should be slow http server")
	return nil
}
