package snip

import (
	"time"
)

// Period represents collection of snippets within a period
type Period struct {

	// PeriodStart is the sunday of the week where period starts
	PeriodStart time.Time `json:"pst"`

	// ObjectiveSnippets is a collection of snippets groups by objective s
	ObjectiveSnippets map[string][]*Snippet `json:"objs"`
}
