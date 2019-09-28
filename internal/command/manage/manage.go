package manage

import (
	"fmt"

	"github.com/urfave/cli"
)

var (
	// AddSnipCommand adds new snippet
	AddSnipCommand = cli.Command{
		Name:     "add",
		Category: "Manage",
		Usage:    "add new snippet",
		Action:   addSnip,
	}
)

func addSnip(c *cli.Context) error {
	fmt.Println("todo - capture new snippet: ", c.Args().First())
	return nil
}
