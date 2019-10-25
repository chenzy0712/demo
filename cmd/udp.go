package cmd

import (
	"github.com/klec/demo/study"
	"github.com/urfave/cli"
)

var (
	UDP = cli.Command{
		Name:        "udp",
		Usage:       "demo udp <option>",
		Description: "run demo UDP case",
		Subcommands: []cli.Command{
			subCmdUDPServer,
			subCmdUDPClient,
		},
	}

	subCmdUDPServer = cli.Command{
		Name:        "server",
		Usage:       "demo udp server",
		Description: "run demo UDP server",
		Flags: []cli.Flag{
			stringFlag("addr, a", "localhost:8888", "UDP server address"),
			intFlag("worker, w", 40, "max workers for data handlers"),
		},
		Action: runUDPServer,
	}

	subCmdUDPClient = cli.Command{
		Name:        "client",
		Usage:       "demo udp client",
		Description: "run demo UDP client",
		Flags: []cli.Flag{
			stringFlag("addr, a", "localhost:8888", "UDP server address"),
			intFlag("interval, i", 5, "UDP client send data interval, us"),
		},
		Action: runUDPClient,
	}
)

func runUDPServer(c *cli.Context) error {
	study.RoutineDemo(c.Int("worker"))
	study.UDPServer(c.String("addr"))

	return nil
}

func runUDPClient(c *cli.Context) error {
	study.UDPClient(c.String("addr"), c.Int("interval"))

	return nil
}
