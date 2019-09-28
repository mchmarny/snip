package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli"
)

var (
	logger = log.New(os.Stdout, "[APP] ", 0)
)

func main() {
	app := cli.NewApp()
	app.Name = "snip"
	app.Version = "v0.1.1"
	app.Usage = "Snippet management utility"

	app.Commands = []cli.Command{
		{
			Name:     "config",
			Category: "Config",
			Usage:    "configuration options",
			Subcommands: []cli.Command{
				{
					Name:  "init",
					Usage: "reinitialize the snip configuration",
					Action: func(c *cli.Context) error {
						fmt.Println("todo - reinitialize configuration ", c.Args().First())
						return nil
					},
				},
			},
		},
		{
			Name:     "add",
			Category: "Manage",
			Usage:    "add new snippet",
			Action: func(c *cli.Context) error {
				fmt.Println("todo - capture new snippet: ", c.Args().First())
				return nil
			},
		},
		{
			Name:     "list",
			Category: "Report",
			Flags: []cli.Flag{
				cli.StringFlag{Name: "period, p"},
			},
			Usage: "lists snippets for specified period",
			Action: func(c *cli.Context) error {
				fmt.Println("todo - list for period:", c.String("period"))
				return nil
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
