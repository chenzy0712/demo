package cmd

import (
	"github.com/urfave/cli"

	"github.com/klec/demo/pkg/log"

	"github.com/klec/demo/internal/httpclient"
)

var (
	Http = cli.Command{
		Name:        "http",
		Usage:       "demo http <option>",
		Description: "run demo test",
		Subcommands: []cli.Command{
			subCmdHttpCli,
		},
	}

	subCmdHttpCli = cli.Command{
		Name:        "client",
		Usage:       "demo http <option>",
		Description: "run client for http",
		Action:      runHttpClient,
	}
)

func runHttpClient(c *cli.Context) error {
	log.Info("Run Http client demo.")
	httpclient.ClientDemo()
	return nil
}
