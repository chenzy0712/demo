package cmd

import (
	"git.kldmp.com/learning/demo/internal/wire/binding"
	"git.kldmp.com/learning/demo/internal/wire/simple"
	"github.com/urfave/cli"
)

var (
	Wire = cli.Command{
		Name:        "wire",
		Usage:       "demo wire <option>",
		Description: "run demo wire",
		Subcommands: []cli.Command{
			subCmdSimpleWire,
			subCmdInterfaceWire,
		},
	}

	subCmdSimpleWire = cli.Command{
		Name:        "simple",
		Usage:       "demo wire simple",
		Description: "simple wire example",
		Action:      runSimpleWire,
	}

	subCmdInterfaceWire = cli.Command{
		Name:        "interface",
		Usage:       "demo wire interface",
		Description: "interface wire example",
		Action:      runInterfaceWire,
	}
)

func runSimpleWire(ctx *cli.Context) error {
	simple.Demo()
	return nil
}

func runInterfaceWire(ctx *cli.Context) error {
	binding.Demo()
	return nil
}
