package main

import (
	"os"

	"github.com/klec/demo/cmd"
	"github.com/klec/demo/pkg/log"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	//app.Name = setting.AppName
	//app.Version = setting.AppVer
	app.Usage = "Golang demo"

	app.Commands = []cli.Command{
		cmd.UT,
	}

	log.Info("Hello demo!")
	err := app.Run(os.Args)
	if err != nil {
		log.Error("%s", err)
	}
}
