package cmd

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/mchmarny/snip/pkg/snip"

	"github.com/urfave/cli"
)

var (
	// ReportCommand lists all snippets for specific period
	ReportCommand = cli.Command{
		Name:     "list",
		Category: "Report",
		Flags: []cli.Flag{
			cli.IntFlag{Name: "week-offset, w"},
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

	weekOffset := c.Int("week-offset")
	weekStart := getWeekPeriodStart(weekOffset)

	wr := os.Stdout
	outputPath := c.String("output")
	if outputPath != "" {
		file, err := os.Create(outputPath)
		if err != nil {
			return fmt.Errorf("error creating output file (%s): %v",
				outputPath, err)
		}
		defer file.Close()
		wr = file
	}

	pr, err := getPeriodSnippets(weekStart)
	if err != nil {
		return fmt.Errorf("error quering data: %v", err)
	}

	fmt.Fprintf(wr, "#Snippets Since: %s\n",
		pr.PeriodStart.Format(snip.SnippetDateFormat))

	for c, s := range pr.ObjectiveSnippets {
		fmt.Fprintf(wr, "\n##%s\n\n", c)
		for _, si := range s {
			fmt.Fprintf(wr, "* %s - %s\n",
				si.CreationTime.Format(snip.SnippetDateFormat),
				si.Text)
		}
	}

	return nil
}

func rankPeriod(c *cli.Context) error {
	return errors.New("not implemented yet")
}

func getWeekPeriodStart(offset int) time.Time {
	now := time.Now()
	today := now.Weekday()
	lastSunday := now.AddDate(0, 0, -int(today))
	periodStart := lastSunday.AddDate(0, 0, -(offset * 7))
	return periodStart
}
