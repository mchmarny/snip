package manage

import (
	"fmt"

	"github.com/mchmarny/snip/pkg/snip"

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
	var list []*snip.Snippet
	fmt.Println("todo - capture new snippet: ", c.Args().First(), len(list))
	return nil
}
