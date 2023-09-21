package main

import (
	"context"
	"os"

	app "github.com/flytrap/go-template/internal"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

var VERSION = "0.0.1"

func main() {
	ctx := context.Background()
	app := cli.NewApp()
	app.Name = "go-template"
	app.Version = VERSION
	app.Usage = "go-template based on GRPC + WIRE."
	app.Commands = []*cli.Command{
		newIndexCmd(ctx),
	}
	err := app.Run(os.Args)
	if err != nil {
		logrus.Error(err.Error())
	}
}

func newIndexCmd(ctx context.Context) *cli.Command {
	return &cli.Command{
		Name:  "run",
		Usage: "Run go template server",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "conf",
				Aliases:  []string{"c"},
				Required: true,
				Usage:    "指定启动配置文件(.json,.yaml,.toml)",
			},
		},
		Action: func(c *cli.Context) error {
			return app.RunApp(ctx,
				app.SetConfigFile(c.String("conf")),
				app.SetVersion(VERSION),
			)
		},
	}
}
