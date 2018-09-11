package connectionfail

import (
	"fmt"
	"net"

	"gitlab.com/xonvanetta/gerish/httpserver"

	"github.com/urfave/cli"
)

//CLICommand the command that runs a server with connection fails
var CLICommand = cli.Command{
	Name:    "connectionfail",
	Aliases: []string{"cf"},
	Usage:   "Starts a server which cuts off connections",
	Action:  action,
	Flags:   cliFlags(),
}

func action(c *cli.Context) {
	connectionfail(c)
}

func connectionfail(context *cli.Context) {

	l, err := net.Listen("tcp", ":"+context.String("port"))
	if err != nil {
		panic(err)
	}
	fmt.Printf("Listening on: %v\n", l.Addr())
	defer l.Close()

	for {
		c, err := l.Accept()

		if err != nil {
			panic(err)
		}

		fmt.Println("Connected")

		go func(c net.Conn) {
			defer c.Close()
			c.(*net.TCPConn).SetLinger(0)
			fmt.Println("Ignoring incoming data")
		}(c)
	}
}

func cliFlags() []cli.Flag {
	return append(
		[]cli.Flag{},
		httpserver.CLIFlags("connectionfail")...,
	)
}
