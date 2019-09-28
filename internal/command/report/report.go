package report

import (
	"fmt"

	"github.com/urfave/cli"
)

var (
	// ReportCommand lists all snippets for specific period
	ReportCommand = cli.Command{
		Name:     "list",
		Category: "Report",
		Flags: []cli.Flag{
			cli.StringFlag{Name: "period, p"},
		},
		Usage:  "lists snippets for specified period",
		Action: reportPeriod,
	}
)

func reportPeriod(c *cli.Context) error {
	fmt.Println("todo - list for period:", c.String("period"))
	return nil
}
