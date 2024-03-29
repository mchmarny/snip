package cmd

import (
	"errors"
	"fmt"
	"log"
	"regexp"
	"strings"
	"time"

	"github.com/mchmarny/snip/pkg/snip"

	"github.com/urfave/cli/v2"
)

const (
	objectiveToken   = "^"
	objectiveRegExp  = `(?:^|\s)\^(\w+)\b`
	contextRegExp    = `(?:^|\s)\@(\w+)\b`
	objectiveDefault = "no-objective"
)

var (
	// AddSnipCommand adds new snippet
	AddSnipCommand = &cli.Command{
		SkipFlagParsing: true,
		Name:            "add",
		Category:        "Manage",
		Usage:           "add new snippet",
		Action:          addSnip,
	}
)

func addSnip(c *cli.Context) error {
	raw := strings.Join(c.Args().Slice(), " ")
	log.Printf("raw: %s", raw)

	// parse
	s, e := parseSnippet(raw)
	if e != nil {
		return fmt.Errorf("error parsing snippet: %v", e)
	}

	// save
	if e = saveSnippet(s); e != nil {
		return fmt.Errorf("error saving item %+s: %v", s, e)
	}

	log.Println("snippet saved")
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

	// context
	ctxList, e := parseItems(text, contextRegExp)
	if e != nil {
		return nil, fmt.Errorf("error parsing context: %v", err)
	}
	s.Contexts = ctxList

	// objectives
	objList, e := parseItems(text, objectiveRegExp)
	if e != nil {
		return nil, fmt.Errorf("error parsing objectives: %v", err)
	}
	if len(objList) != 1 {
		objList = []string{objectiveDefault}
	}

	s.Objective = strings.TrimSpace(strings.ReplaceAll(objList[0], objectiveToken, ""))

	// text, replace all objectives
	txt := s.Raw
	for _, o := range objList {
		txt = strings.ReplaceAll(txt, o, "")
	}
	s.Text = strings.TrimSpace(txt)

	return s, nil
}

func parseItems(s, exp string) (items []string, err error) {
	r, e := regexp.Compile(exp)
	if e != nil {
		return nil, e
	}
	parts := r.FindAllString(s, -1) // nil on no match
	if parts == nil {               // parseItems always returns empty array s
		parts = []string{}
	}

	list := make([]string, len(parts))
	for i, p := range parts {
		list[i] = strings.ReplaceAll(p, " ", "")
	}

	return list, nil
}
