package snip

import (
	"fmt"
	"strings"
	"time"
)

const (
	// SnippetDateFormat is "2006-01-02"
	SnippetDateFormat = "2006-01-02"
	// SnippetDateTimeFormat is "2006-01-02 15:04"
	SnippetDateTimeFormat = "2006-01-02 15:04"
)

// Snippet represents a single snippet
type Snippet struct {
	// Raw is the original string that was entered by the user
	Raw string `json:"raw"`
	// Text is the parsed snippet, raw sans objective
	Text string `json:"text"`
	// CreationTime is the time when the snippet was created
	CreationTime time.Time `json:"ctm"`
	// Objective is the extracted #objective from raw
	Objective string `json:"obj"`
	// Contexts is the extracted @context from raw
	Contexts []string `json:"ctx"`
}

func (s Snippet) String() string {
	return fmt.Sprintf("%s (text:%s \non:%s \nobjective:%s \ncontexts[%d]:%s)",
		s.Raw,
		s.Text,
		s.CreationTime.Format(SnippetDateTimeFormat),
		s.Objective,
		len(s.Contexts),
		strings.Join(s.Contexts, ","))
}
