// Package agents promotes a Claude specific instruction file to the portable
// AGENTS.md convention.
//
// The content `claude /init` writes is useful to every coding agent, only its
// name and a few boilerplate phrases are Claude specific. Migrate moves the
// file to AGENTS.md and leaves pointers behind for the tools looking for their
// own filename.
package agents

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/charmbracelet/log"
	"github.com/spf13/afero"
)

const (
	AgentsFile  = "AGENTS.md"
	ClaudeFile  = "CLAUDE.md"
	CopilotFile = ".github/copilot-instructions.md"
)

// claudePointer imports AGENTS.md with Claude Code's `@` syntax so the content
// is pulled in rather than merely linked
const claudePointer = `# ` + ClaudeFile + `

@` + AgentsFile + `
`

const copilotPointer = `# Copilot instructions

See [` + AgentsFile + `](../` + AgentsFile + `).
`

type Options struct {
	Force bool
}

// Migrate promotes CLAUDE.md to AGENTS.md and leaves pointers behind.
//
// CLAUDE.md is always rewritten, it is the source of the migration. AGENTS.md
// and the copilot instructions are only overwritten when Force is set.
//
// The returned references are the lines of AGENTS.md still mentioning Claude,
// they are reported rather than guessed at.
func Migrate(fs afero.Fs, options Options) ([]Reference, error) {
	content, err := afero.ReadFile(fs, ClaudeFile)
	if os.IsNotExist(err) {
		return nil, fmt.Errorf("no %s to migrate", ClaudeFile)
	} else if err != nil {
		return nil, err
	}

	// a second run would otherwise promote the pointer over the very
	// instructions it points at, --force included
	if imports(string(content)) {
		return nil, fmt.Errorf("%s already imports %s, nothing to migrate",
			ClaudeFile, AgentsFile,
		)
	}

	if !options.Force {
		for _, path := range []string{AgentsFile, CopilotFile} {
			if exists, err := afero.Exists(fs, path); err != nil {
				return nil, err
			} else if exists {
				return nil, fmt.Errorf("%s already exists, pass --force to overwrite", path)
			}
		}
	}

	agents, remaining := Transform(string(content))
	if err = write(fs, AgentsFile, agents); err != nil {
		return nil, err
	}

	if err = write(fs, ClaudeFile, claudePointer); err != nil {
		return nil, err
	}

	if err = fs.MkdirAll(filepath.Dir(CopilotFile), os.ModePerm); err != nil {
		return nil, err
	}
	if err = write(fs, CopilotFile, copilotPointer); err != nil {
		return nil, err
	}

	return remaining, nil
}

// imports reports whether content already pulls in AGENTS.md
func imports(content string) bool {
	for line := range strings.Lines(content) {
		if strings.TrimSpace(line) == "@"+AgentsFile {
			return true
		}
	}

	return false
}

func write(fs afero.Fs, path, content string) error {
	if err := afero.WriteFile(fs, path, []byte(content), 0o644); err != nil {
		return err
	}

	log.Debugf("wrote %s", path)
	return nil
}
