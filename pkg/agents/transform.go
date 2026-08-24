package agents

import (
	"regexp"
	"strings"
)

// A Reference is a line of AGENTS.md still mentioning Claude after the
// boilerplate rewrites
type Reference struct {
	Line int
	Text string
}

var (
	titleExpr    = regexp.MustCompile(`^#\s+CLAUDE\.md\s*$`)
	guidanceExpr = regexp.MustCompile(`Claude Code \(claude\.ai/code\)`)
	selfExpr     = regexp.MustCompile(`\bCLAUDE\.md\b`)
	claudeExpr   = regexp.MustCompile(`(?i)claude`)
)

// Transform rewrites the boilerplate `claude /init` emits and returns the
// result along with every line still mentioning Claude.
//
// Only the fixed phrases are rewritten. Anything else is left alone and
// reported so it can be edited by hand.
func Transform(content string) (string, []Reference) {
	lines := strings.Split(content, "\n")

	titled := false
	for i, line := range lines {
		if !titled && titleExpr.MatchString(line) {
			line, titled = "# "+AgentsFile, true
		}

		line = guidanceExpr.ReplaceAllString(line, "coding agents")
		lines[i] = selfExpr.ReplaceAllString(line, AgentsFile)
	}

	var remaining []Reference
	for i, line := range lines {
		if claudeExpr.MatchString(line) {
			remaining = append(remaining, Reference{
				Line: i + 1,
				Text: strings.TrimSpace(line),
			})
		}
	}

	return strings.Join(lines, "\n"), remaining
}
