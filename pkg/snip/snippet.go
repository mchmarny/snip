package snip

import (
	"fmt"
	"strings"
	"time"
)

// Snippet represents a single snippet
type Snippet struct {
	// Raw is the original string that was entered by the user
	Raw string
	// Text is the parsed snippet text
	Text string
	// CreationTime is the time when the snippet was created
	CreationTime time.Time
	// Contexts is the extracted @context from raw
	Contexts []string
	// Objectives is the extracted #objective from raw
	Objectives []string
}

func (s Snippet) String() string {
	return fmt.Sprintf("%s (on:%s objectives[%d]:%s contexts[%d]:%s)",
		s.Raw,
		s.CreationTime.Format("2006-01-02 15:04"),
		len(s.Objectives),
		strings.Join(s.Objectives, ","),
		len(s.Contexts),
		strings.Join(s.Contexts, ","))
}
