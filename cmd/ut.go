package cmd

import (
	"git.kldmp.com/learning/demo/internal/model"
	"git.kldmp.com/learning/demo/internal/po"
	"git.kldmp.com/learning/demo/internal/po/db"
	"git.kldmp.com/learning/demo/study"
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
