package e2e_test

import (
	"os"
	"os/exec"
	"path/filepath"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"
	"github.com/onsi/gomega/gexec"
	"github.com/unmango/devctl/pkg/agents"
)

var _ = Describe("migrate", func() {
	var root string

	devctl := func(args ...string) *gexec.Session {
		cmd := exec.Command(cmdPath, args...)
		cmd.Dir = root

		ses, err := gexec.Start(cmd, GinkgoWriter, GinkgoWriter)
		Expect(err).NotTo(HaveOccurred())
		return ses
	}

	read := func(path string) string {
		data, err := os.ReadFile(filepath.Join(root, path))
		Expect(err).NotTo(HaveOccurred())
		return string(data)
	}

	BeforeEach(func() {
		root = GinkgoT().TempDir()
	})

	Describe("agents", func() {
		BeforeEach(func() {
			Expect(os.WriteFile(
				filepath.Join(root, agents.ClaudeFile),
				[]byte("# CLAUDE.md\n\nRun `make build`. Update CLAUDE.md when it changes.\n"),
				os.ModePerm,
			)).To(Succeed())
		})

		It("should migrate to the AGENTS.md convention", func() {
			ses := devctl("migrate", "agents", "-C", root)

			Eventually(ses).Should(gexec.Exit(0))
			Expect(read(agents.AgentsFile)).To(Equal(
				"# AGENTS.md\n\nRun `make build`. Update AGENTS.md when it changes.\n",
			))
			Expect(read(agents.ClaudeFile)).To(Equal("# CLAUDE.md\n\n@AGENTS.md\n"))
			Expect(read(agents.CopilotFile)).To(ContainSubstring("../AGENTS.md"))
		})

		It("should report the lines it left alone", func() {
			Expect(os.WriteFile(
				filepath.Join(root, agents.ClaudeFile),
				[]byte("# CLAUDE.md\n\nAsk Claude to run the tests.\n"),
				os.ModePerm,
			)).To(Succeed())

			ses := devctl("migrate", "agents", "-C", root)

			Eventually(ses).Should(gexec.Exit(0))
			Expect(ses.Err).To(gbytes.Say(
				`AGENTS.md:3 still mentions Claude: Ask Claude to run the tests.`,
			))
		})

		It("should refuse to overwrite an existing AGENTS.md", func() {
			Expect(os.WriteFile(
				filepath.Join(root, agents.AgentsFile),
				[]byte("Mine"),
				os.ModePerm,
			)).To(Succeed())

			ses := devctl("migrate", "agents", "-C", root)

			Eventually(ses).ShouldNot(gexec.Exit(0))
			Expect(read(agents.AgentsFile)).To(Equal("Mine"))
		})

		It("should overwrite an existing AGENTS.md when forced", func() {
			Expect(os.WriteFile(
				filepath.Join(root, agents.AgentsFile),
				[]byte("Mine"),
				os.ModePerm,
			)).To(Succeed())

			ses := devctl("migrate", "agents", "-C", root, "--force")

			Eventually(ses).Should(gexec.Exit(0))
			Expect(read(agents.AgentsFile)).To(ContainSubstring("# AGENTS.md"))
		})
	})
})
