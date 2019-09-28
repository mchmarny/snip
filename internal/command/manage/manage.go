package manage

import (
	"errors"
	"fmt"
	"log"
	"regexp"
	"strings"
	"time"

	"github.com/mchmarny/snip/pkg/snip"

	"github.com/urfave/cli"
)

const (
	objectiveToken  = "^"
	objectiveRegExp = `(?:^|\s)\^(\w+)\b`
	contextToken    = "@"
	contextRegExp   = `(?:^|\s)\@(\w+)\b`
)

var (
	// AddSnipCommand adds new snippet
	AddSnipCommand = cli.Command{
		SkipFlagParsing: true,
		SkipArgReorder:  true,
		Name:            "add",
		Category:        "Manage",
		Usage:           "add new snippet",
		Action:          addSnip,
	}
)

func addSnip(c *cli.Context) error {

	raw := strings.Join([]string(c.Args()), " ")
	log.Printf("raw: %s", raw)

	s, e := parseSnippet(raw)

	if e != nil {
		return e
	}

	fmt.Printf("snip: %s\n", s.String())

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
	list, e := parseItems(text, objectiveToken, objectiveRegExp)
	if e != nil {
		return nil, fmt.Errorf("error parsing objectives: %v", err)
	}
	log.Printf("found %d objectives", len(list))
	s.Objectives = list

	// context
	list, e = parseItems(text, contextToken, contextRegExp)
	if e != nil {
		return nil, fmt.Errorf("error parsing context: %v", err)
	}
	log.Printf("found %d contexts", len(list))
	s.Contexts = list

	return s, nil
}

func parseItems(s, t, exp string) (items []string, err error) {
	r, e := regexp.Compile(exp)
	if e != nil {
		return nil, e
	}
	parts := r.FindAllString(s, -1) // nil on no match
	if parts == nil {               // parseItems always returns empty array s
		parts = []string{}
	}

	for i, p := range parts {
		// trim spaces and then the token
		parts[i] = strings.Trim(strings.Trim(p, " "), t)
	}

	return parts, nil
}
