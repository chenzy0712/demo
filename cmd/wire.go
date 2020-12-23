package cmd

import (
	"github.com/klec/demo/internal/wire/simple"
	"github.com/urfave/cli"
)

var (
	Wire = cli.Command{
		Name:        "wire",
		Usage:       "demo wire <option>",
		Description: "run demo wire",
		Subcommands: []cli.Command{
			subCmdSimpleWire,
		},
	}

	subCmdSimpleWire = cli.Command{
		Name:        "simple",
		Usage:       "demo wire simple",
		Description: "simple wire example",
		Action:      runSimpleWire,
	}
)

func runSimpleWire(ctx *cli.Context) error {
	simple.Demo()
	return nil
}
