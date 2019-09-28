package manage

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/mchmarny/snip/pkg/snip"

	"github.com/urfave/cli"
)

const (
	tagRegExp = `(?:^|\s)\#(\w+)\b`
	ctxRegExp = `(?:^|\s)\@(\w+)\b`
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

	// TODO: find cleaner way to parse the entire command
	s, e := parseSnippet(fmt.Sprintf("%s %s",
		c.Args().First(), strings.Join(c.Args().Tail(), " ")))

	if e != nil {
		return e
	}

	fmt.Printf("snip: %s", s.String())

	return nil
}

func parseSnippet(text string) (snippet *snip.Snippet, err error) {
	if text == "" {
		return nil, errors.New("text required")
	}

	s := &snip.Snippet{
		Raw:          text,
		CreationTime: time.Now(),
	}

	// tags
	list, e := parseItems(text, tagRegExp)
	if e != nil {
		return nil, fmt.Errorf("error parsing tags: %v", err)
	}
	s.Tags = list

	// context
	list, e = parseItems(text, ctxRegExp)
	if e != nil {
		return nil, fmt.Errorf("error parsing context: %v", err)
	}
	s.Contexts = list

	return s, nil
}

func parseItems(s, exp string) (items []string, err error) {
	r, e := regexp.Compile(exp)
	if e != nil {
		return nil, e
	}
	list := r.FindAllString(s, -1)
	return list, nil
}
