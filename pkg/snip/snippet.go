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
	// ID is a hash of the creation time
	ID string
	// Raw is the original string that was entered by the user
	Raw string
	// Text is the parsed snippet, raw sans objective
	Text string
	// CreationTime is the time when the snippet was created
	CreationTime time.Time
	// Objective is the extracted #objective from raw
	Objective string
	// Contexts is the extracted @context from raw
	Contexts []string
}

func (s Snippet) String() string {
	return fmt.Sprintf("%s (id:%s \ntext:%s \non:%s \nobjective:%s \ncontexts[%d]:%s)",
		s.ID,
		s.Raw,
		s.Text,
		s.CreationTime.Format(SnippetDateTimeFormat),
		s.Objective,
		len(s.Contexts),
		strings.Join(s.Contexts, ","))
}