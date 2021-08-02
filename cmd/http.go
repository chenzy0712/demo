package cmd

import (
	"context"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	transhttp "github.com/go-kratos/kratos/v2/transport/http"
	"github.com/urfave/cli"

	"git.kldmp.com/learning/demo/pkg/log"

	"git.kldmp.com/learning/demo/bizapi"
	"git.kldmp.com/learning/demo/internal/http/client"
	service "git.kldmp.com/learning/demo/internal/http/server"
)

var (
	Http = cli.Command{
		Name:        "http",
		Usage:       "demo http <option>",
		Description: "run demo test",
		Subcommands: []cli.Command{
			subCmdEmqxHttpCli,
			subCmdKratosHttpCli,
			subCmdKratosHttpSrv,
		},
	}

	subCmdEmqxHttpCli = cli.Command{
		Name:        "client",
		Usage:       "demo http <option>",
		Description: "run client for http",
		Action:      runHttpClient,
	}

	subCmdKratosHttpCli = cli.Command{
		Name:        "kclient",
		Usage:       "demo http kclient <option>",
		Description: "run http client for kratos",
		Action:      runKratosHttpClient,
	}

	subCmdKratosHttpSrv = cli.Command{
		Name:        "kserver",
		Usage:       "demo http kserver <option>",
		Description: "run http server for kratos",
		Action:      runKratosHttpServer,
	}
)

func runHttpClient(c *cli.Context) error {
	log.Info("Run Http client demo.")
	client.ClientDemo()
	return nil
}

func runKratosHttpClient(c *cli.Context) error {
	conn, err := transhttp.NewClient(
		context.Background(),
		transhttp.WithMiddleware(
			recovery.Recovery(),
		),
		transhttp.WithEndpoint("127.0.0.1:8000"),
	)
	if err != nil {
		panic(err)
	}

	bizClient := bizapi.NewBizHTTPClient(conn)
	reply, err := bizClient.GetDataTransferProtocolList(context.Background(), &bizapi.GetDataTransferProtocolListReq{ProductKey: "IEC104s"})
	if err != nil {
		log.Error("%s", err)
	}
	log.Info("[http] GetDataTransferProtocolList %+v", reply.Protocols)

	return nil
}

func runKratosHttpServer(c *cli.Context) error {
	log.Info("Start biz http server running...")

	httpSrv := transhttp.NewServer(
		transhttp.Address(":8000"),
		transhttp.Middleware(
			recovery.Recovery(),
		),
	)

	s := service.NewBizService()
	bizapi.RegisterBizHTTPServer(httpSrv, s)

	app := kratos.New(
		kratos.Name("bizapi"),
		kratos.Server(
			httpSrv,
		),
	)

	if err := app.Run(); err != nil {
		log.Error("App running error:%s", err)
	}

	return nil
}
