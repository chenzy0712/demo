package main

import (
	"os"

	"git.kldmp.com/learning/demo/cmd"
	"git.kldmp.com/learning/demo/pkg/log"
	"git.kldmp.com/learning/demo/pkg/setting"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = setting.AppName
	app.Version = setting.AppVer
	app.Usage = "Golang demo"

	app.Commands = []cli.Command{
		cmd.UT,
		cmd.UDP,
		cmd.CRC,
		cmd.TSDB,
		cmd.Ctx,
		cmd.Wire,
		cmd.Linear,
		cmd.Test,
		cmd.Http,
	}

	log.InitLog("console", "", "info", 10)
	err := app.Run(os.Args)
	if err != nil {
		log.Error("%s", err)
	}
}
