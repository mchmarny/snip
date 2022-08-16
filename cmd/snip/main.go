package main

import (
	"fmt"
	"os"

	"time"

	"github.com/mchmarny/snip/internal/cmd"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

var (
	name    = "snip"
	version = "v0.0.1-default"
	commit  = ""
	date    = ""

	debug = false

	debugFlag = &cli.BoolFlag{
		Name:        "debug",
		Usage:       "Prints verbose logs (optional, default: false)",
		Destination: &debug,
	}
)

func main() {
	initLogging(name, version)

	var err error
	if err = cmd.Init(); err != nil {
		fatalErr(err)
	}

	app := &cli.App{
		Name:     "snip",
		Version:  fmt.Sprintf("%s (%s - %s)", version, commit, date),
		Compiled: time.Now(),
		Usage:    "Simple utility to collect snippets.",
		Flags: []cli.Flag{
			debugFlag,
		},
		Commands: []*cli.Command{
			cmd.AddSnipCommand,
			cmd.ReportCommand,
		},
		Before: func(c *cli.Context) error {
			if c.Bool(debugFlag.Name) {
				log.SetLevel(log.DebugLevel)
				// log.SetReportCaller(true)
			}
			return nil
		},
	}

	err = app.Run(os.Args)
	if err != nil {
		fatalErr(err)
	}
}

func fatalErr(err error) {
	if err != nil {
		log.Fatalf("fatal error: %v", err)
		os.Exit(1)
	}
}

func initLogging(name, version string) {
	log.SetOutput(os.Stdout)
	log.SetLevel(log.InfoLevel)
	log.SetReportCaller(false)
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp:          false,
		DisableTimestamp:       true,
		ForceColors:            true,
		DisableLevelTruncation: true,
		PadLevelText:           true,
	})
}
