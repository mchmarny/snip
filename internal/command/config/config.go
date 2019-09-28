package config

import (
	"fmt"

	"github.com/urfave/cli"
)

var (
	// InitConfigCommand re-initializes app config
	InitConfigCommand = cli.Command{
		Name:     "config",
		Category: "Config",
		Usage:    "configuration options",
		Subcommands: []cli.Command{
			{
				Name:   "init",
				Usage:  "reinitialize the snip configuration",
				Action: initConfig,
			},
		},
	}
)

func initConfig(c *cli.Context) error {
	fmt.Println("todo - reinitialize configuration ", c.Args().First())
	return nil
}
