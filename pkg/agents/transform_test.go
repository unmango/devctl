package agents_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/unmango/devctl/pkg/agents"
)

var _ = Describe("Transform", func() {
	It("should rewrite the title", func() {
		content, remaining := agents.Transform("# CLAUDE.md\n\nStuff\n")

		Expect(content).To(Equal("# AGENTS.md\n\nStuff\n"))
		Expect(remaining).To(BeEmpty())
	})

	It("should rewrite the guidance sentence", func() {
		content, remaining := agents.Transform(
			"This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.",
		)

		Expect(content).To(Equal(
			"This file provides guidance to coding agents when working with code in this repository.",
		))
		Expect(remaining).To(BeEmpty())
	})

	It("should rewrite self references", func() {
		content, _ := agents.Transform("Update CLAUDE.md when the commands change.")

		Expect(content).To(Equal("Update AGENTS.md when the commands change."))
	})

	It("should only rewrite the first title", func() {
		content, _ := agents.Transform("# CLAUDE.md\n\n# CLAUDE.md\n")

		Expect(content).To(Equal("# AGENTS.md\n\n# AGENTS.md\n"))
	})

	It("should report lines it left alone", func() {
		content, remaining := agents.Transform(
			"# CLAUDE.md\n\nAsk Claude to run the tests.\n\nRun `make build` with Claude Code.\n",
		)

		Expect(content).To(Equal(
			"# AGENTS.md\n\nAsk Claude to run the tests.\n\nRun `make build` with Claude Code.\n",
		))
		Expect(remaining).To(Equal([]agents.Reference{
			{Line: 3, Text: "Ask Claude to run the tests."},
			{Line: 5, Text: "Run `make build` with Claude Code."},
		}))
	})

	It("should leave unrelated content alone", func() {
		content, remaining := agents.Transform("# Build\n\nRun `make build`.\n")

		Expect(content).To(Equal("# Build\n\nRun `make build`.\n"))
		Expect(remaining).To(BeEmpty())
	})
})
