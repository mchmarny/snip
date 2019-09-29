package main

import (
	"log"
	"os"

	"github.com/mchmarny/snip/internal/cmd"

	"github.com/urfave/cli"
)

const (
	appName    = "snip"
	appVersion = "v0.1.1"
)

func main() {

	log.SetOutput(os.Stdout)
	log.SetFlags(0)

	app := cli.NewApp()
	app.Name = appName
	app.Version = appVersion
	app.Usage = "Snippet management utility"
	app.Commands = []cli.Command{
		cmd.InitConfigCommand,
		cmd.AddSnipCommand,
		cmd.ReportCommand,
		cmd.RankCommand,
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
