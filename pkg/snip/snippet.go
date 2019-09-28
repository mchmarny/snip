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
	// Tags is the extracted #tags from raw
	Tags []string
}

func (s Snippet) String() string {
	return fmt.Sprintf("Raw:%s On:%s Tags:%s Ctx:%s",
		s.Raw, s.CreationTime.String(),
		strings.Join(s.Tags, ","),
		strings.Join(s.Contexts, ","))
}
