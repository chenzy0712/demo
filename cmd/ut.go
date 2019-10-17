package cmd

import (
	"github.com/klec/demo/internal/model"
	"github.com/klec/demo/internal/po"
	"github.com/klec/demo/internal/po/db"
	"github.com/klec/demo/study"
	"github.com/urfave/cli"
)

var (
	UT = cli.Command{
		Name:        "ut",
		Usage:       "demo ut <option>",
		Description: "run demo UT case",
		Subcommands: []cli.Command{
			subCmdDemoAll,
		},
	}

	subCmdDemoAll = cli.Command{
		Name:        "all",
		Usage:       "demo ut all",
		Description: "run all UT case",
		Flags:       []cli.Flag{},
		Action:      runAll,
	}
)

func runAll(c *cli.Context) error {
	po.SetPo(db.NewXormPo())

	model.Demo()

	study.InterfaceDemo()

	return nil
}
