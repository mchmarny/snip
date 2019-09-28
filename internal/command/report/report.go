package report

import (
	"fmt"

	"github.com/mchmarny/snip/pkg/snip"

	"github.com/urfave/cli"
)

var (
	// ReportCommand lists all snippets for specific period
	ReportCommand = cli.Command{
		Name:     "list",
		Category: "Report",
		Flags: []cli.Flag{
			cli.StringFlag{Name: "period, p"},
			cli.StringFlag{Name: "output, o"},
		},
		Usage:  "lists snippets for specified period",
		Action: reportPeriod,
	}

	// RankCommand ranks snippets based on number of tags for specific period
	RankCommand = cli.Command{
		Name:     "rank",
		Category: "Report",
		Flags: []cli.Flag{
			cli.StringFlag{Name: "period, p"},
		},
		Usage:  "lists snippets based on tags for specified period",
		Action: rankPeriod,
	}
)

func reportPeriod(c *cli.Context) error {
	var list []*snip.Snippet
	fmt.Println("todo - list for period:", c.String("period"), len(list))
	return nil
}

func rankPeriod(c *cli.Context) error {
	var list []*snip.Snippet
	fmt.Println("todo - list for period:", c.String("period"), len(list))
	return nil
}
