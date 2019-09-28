package snip

import "time"

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
