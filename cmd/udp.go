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
			subCmdMulticastServer,
			subCmdMulticastClient,
			subCmdUDPServer,
			subCmdUDPClient,
		},
	}

	subCmdMulticastServer = cli.Command{
		Name:        "multi",
		Usage:       "demo udp multi",
		Description: "run demo UDP multicast server",
		Flags: []cli.Flag{
			stringFlag("addr, a", "239.0.0.0:9999", "UDP server address"),
			intFlag("worker, w", 40, "max workers for data handlers"),
		},
		Action: runMulticastServer,
	}

	subCmdMulticastClient = cli.Command{
		Name:        "multic",
		Usage:       "demo udp multic",
		Description: "run demo UDP multicast client",
		Flags: []cli.Flag{
			stringFlag("addr, a", "239.0.0.0:9999", "UDP server address"),
			intFlag("interval, i", 100, "UDP client send data interval, us"),
		},
		Action: runMulticastClient,
	}

	subCmdUDPServer = cli.Command{
		Name:        "server",
		Usage:       "demo udp server",
		Description: "run demo UDP server",
		Flags: []cli.Flag{
			stringFlag("addr, a", "127.0.0.1:8888", "UDP server address"),
			intFlag("worker, w", 40, "max workers for data handlers"),
		},
		Action: runUDPServer,
	}

	subCmdUDPClient = cli.Command{
		Name:        "client",
		Usage:       "demo udp client",
		Description: "run demo UDP client",
		Flags: []cli.Flag{
			stringFlag("addr, a", "127.0.0.1:8888", "UDP server address"),
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

func runMulticastServer(c *cli.Context) error {
	study.RoutineDemo(c.Int("worker"))
	study.MulticastServer(c.String("addr"))

	return nil
}

func runMulticastClient(c *cli.Context) error {
	study.MulticastClient(c.String("addr"), c.Int("interval"))

	return nil
}
