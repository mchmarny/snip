package main

import (
	"log"
	"os"

	"github.com/mchmarny/snip/internal/command/config"
	"github.com/mchmarny/snip/internal/command/manage"
	"github.com/mchmarny/snip/internal/command/report"

	"github.com/urfave/cli"
)

const (
	appName    = "snip"
	appVersion = "v0.1.1"
)

func main() {
	app := cli.NewApp()
	app.Name = appName
	app.Version = appVersion
	app.Usage = "Snippet management utility"

	app.Commands = []cli.Command{
		config.InitConfigCommand,
		manage.AddSnipCommand,
		report.ReportCommand,
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
