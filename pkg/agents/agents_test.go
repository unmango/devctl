package agents_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/spf13/afero"
	"github.com/unmango/devctl/pkg/agents"
)

// claudeInit is what `claude /init` writes, trimmed to the parts that matter
const claudeInit = `# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Build

Run ` + "`make build`" + `. Update CLAUDE.md when the commands change.
`

var _ = Describe("Migrate", func() {
	var fs afero.Fs

	read := func(path string) string {
		data, err := afero.ReadFile(fs, path)
		Expect(err).NotTo(HaveOccurred())
		return string(data)
	}

	BeforeEach(func() {
		fs = afero.NewMemMapFs()
	})

	Context("with a CLAUDE.md", func() {
		BeforeEach(func() {
			Expect(afero.WriteFile(fs,
				agents.ClaudeFile, []byte(claudeInit), 0o644,
			)).To(Succeed())
		})

		It("should write the instructions to AGENTS.md", func() {
			_, err := agents.Migrate(fs, agents.Options{})

			Expect(err).NotTo(HaveOccurred())
			Expect(read(agents.AgentsFile)).To(Equal(`# AGENTS.md

This file provides guidance to coding agents when working with code in this repository.

## Build

Run ` + "`make build`" + `. Update AGENTS.md when the commands change.
`))
		})

		It("should point CLAUDE.md at AGENTS.md", func() {
			_, err := agents.Migrate(fs, agents.Options{})

			Expect(err).NotTo(HaveOccurred())
			Expect(read(agents.ClaudeFile)).To(Equal("@AGENTS.md\n"))
		})

		It("should point the copilot instructions at AGENTS.md", func() {
			_, err := agents.Migrate(fs, agents.Options{})

			Expect(err).NotTo(HaveOccurred())
			Expect(read(agents.CopilotFile)).To(Equal(
				"@../AGENTS.md\n",
			))
		})

		It("should report the lines it left alone", func() {
			Expect(afero.WriteFile(fs,
				agents.ClaudeFile,
				[]byte(claudeInit+"\nAsk Claude to run the tests.\n"),
				0o644,
			)).To(Succeed())

			remaining, err := agents.Migrate(fs, agents.Options{})

			Expect(err).NotTo(HaveOccurred())
			Expect(remaining).To(Equal([]agents.Reference{
				{Line: 9, Text: "Ask Claude to run the tests."},
			}))
		})

		It("should fail when AGENTS.md exists", func() {
			Expect(afero.WriteFile(fs,
				agents.AgentsFile, []byte("Mine"), 0o644,
			)).To(Succeed())

			_, err := agents.Migrate(fs, agents.Options{})

			Expect(err).To(MatchError(ContainSubstring("AGENTS.md already exists")))
			Expect(read(agents.AgentsFile)).To(Equal("Mine"))
			Expect(read(agents.ClaudeFile)).To(Equal(claudeInit))
		})

		It("should fail when the copilot instructions exist", func() {
			Expect(afero.WriteFile(fs,
				agents.CopilotFile, []byte("Mine"), 0o644,
			)).To(Succeed())

			_, err := agents.Migrate(fs, agents.Options{})

			Expect(err).To(MatchError(ContainSubstring("copilot-instructions.md already exists")))
			Expect(read(agents.CopilotFile)).To(Equal("Mine"))
		})

		It("should overwrite existing files when forced", func() {
			Expect(afero.WriteFile(fs,
				agents.AgentsFile, []byte("Mine"), 0o644,
			)).To(Succeed())
			Expect(afero.WriteFile(fs,
				agents.CopilotFile, []byte("Mine"), 0o644,
			)).To(Succeed())

			_, err := agents.Migrate(fs, agents.Options{Force: true})

			Expect(err).NotTo(HaveOccurred())
			Expect(read(agents.AgentsFile)).To(ContainSubstring("# AGENTS.md"))
			Expect(read(agents.CopilotFile)).To(ContainSubstring("../AGENTS.md"))
		})
	})

	It("should fail when CLAUDE.md already imports AGENTS.md", func() {
		Expect(afero.WriteFile(fs,
			agents.ClaudeFile, []byte("# CLAUDE.md\n\n@AGENTS.md\n"), 0o644,
		)).To(Succeed())
		Expect(afero.WriteFile(fs,
			agents.AgentsFile, []byte(claudeInit), 0o644,
		)).To(Succeed())

		_, err := agents.Migrate(fs, agents.Options{Force: true})

		Expect(err).To(MatchError("CLAUDE.md already imports AGENTS.md, nothing to migrate"))
		Expect(read(agents.AgentsFile)).To(Equal(claudeInit))
	})

	It("should fail when there is no CLAUDE.md", func() {
		_, err := agents.Migrate(fs, agents.Options{})

		Expect(err).To(MatchError("no CLAUDE.md to migrate"))
	})
})
