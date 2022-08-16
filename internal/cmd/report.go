package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/mchmarny/snip/pkg/snip"

	"github.com/urfave/cli/v2"
)

const (
	weekDays = 7
)

var (
	offsetFlag = &cli.IntFlag{
		Name:    "week-offset",
		Value:   1,
		Aliases: []string{"w"},
	}

	outputFlag = &cli.StringFlag{
		Name:    "output",
		Aliases: []string{"o"},
	}

	// ReportCommand lists all snippets for specific period
	ReportCommand = &cli.Command{
		Name:     "list",
		Category: "Report",
		Flags: []cli.Flag{
			offsetFlag,
			outputFlag,
		},
		Usage:  "lists snippets for specified period (default: 1)",
		Action: reportPeriod,
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

	fmt.Fprintf(wr, "# Snippets Since: %s\n",
		pr.PeriodStart.Format(snip.SnippetDateFormat))

	for c, s := range pr.ObjectiveSnippets {
		fmt.Fprintf(wr, "\n## %s\n\n", c)
		for _, si := range s {
			fmt.Fprintf(wr, "* %s - %s\n",
				si.CreationTime.Format(snip.SnippetDateFormat),
				si.Text)
		}
	}

	return nil
}

func getWeekPeriodStart(offset int) time.Time {
	now := time.Now()
	today := now.Weekday()
	lastSun := now.AddDate(0, 0, -int(today))
	offsetSun := lastSun.AddDate(0, 0, -(offset * weekDays))
	wkStart := time.Date(offsetSun.Year(), offsetSun.Month(), offsetSun.Day(), 0, 0, 0, 0, time.UTC)
	return wkStart
}
