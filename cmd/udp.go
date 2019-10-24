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
		},
	}

	subCmdUDPServer = cli.Command{
		Name:        "server",
		Usage:       "demo ut server",
		Description: "run demo UDP server",
		Flags: []cli.Flag{
			stringFlag("addr, a", "192.168.12.16:8888", "UDP server address"),
			intFlag("worker, w", 4, "max workers for data handlers"),
		},
		Action: runUDPServer,
	}
)

func runUDPServer(c *cli.Context) error {

	study.RoutineDemo(c.Int("worker"))

	study.UDPServer(c.String("addr"))

	return nil
}
