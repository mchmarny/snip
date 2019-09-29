package cmd

import (
	"errors"
	"fmt"
	"log"
	"strings"
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

	weekStart := getWeekPeriodStart(1)

	list, err := getWeekSnippets(weekStart)
	if err != nil {
		return fmt.Errorf("error quering data: %v", err)
	}

	log.Printf("snippets for week startign with: %s ",
		weekStart.Format(snip.SnippetDateTimeFormat))

	for i, s := range list {
		log.Printf("[%d] id:%s text:%s on:%s objectives:%s contexts:%s",
			i,
			s.ID,
			s.Text,
			s.CreationTime.Format(snip.SnippetDateTimeFormat),
			strings.Join(s.Objectives, ","),
			strings.Join(s.Contexts, ","))
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
