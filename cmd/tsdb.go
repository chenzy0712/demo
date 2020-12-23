package cmd

import (
	"github.com/urfave/cli"
)

var (
	TSDB = cli.Command{
		Name:        "tsdb",
		Usage:       "demo tsdb <option>",
		Description: "run demo TSDB case",
		Subcommands: []cli.Command{
			subCmdTsdb,
		},
	}

	subCmdTsdb = cli.Command{
		Name:        "run",
		Usage:       "demo tsdb run",
		Description: "run TSDB demo",
		Flags: []cli.Flag{
			stringFlag("addr, a", "127.0.0.1:8888", "UDP server address"),
			intFlag("interval, i", 5, "UDP client send data interval, us"),
		},
		Action: runTsdb,
	}
)

func runTsdb(c *cli.Context) error {

	return nil
}
