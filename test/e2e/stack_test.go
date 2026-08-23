package e2e_test

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"
	"github.com/onsi/gomega/gexec"
)

var _ = Describe("stack", func() {
	var root, calls string

	// stub writes an executable recording its invocation to the calls file,
	// exiting with code when its arguments match args exactly.
	stub := func(name, args string, code int) string {
		path := filepath.Join(root, name)
		script := fmt.Sprintf(`#!/bin/sh
printf '%s %%s\n' "$*" >> %q
case "$*" in
%q)
	exit %d
	;;
esac
exit 0
`, name, calls, args, code)

		Expect(os.WriteFile(path, []byte(script), 0o755)).To(Succeed())
		return path
	}

	devctl := func(args ...string) *gexec.Session {
		cmd := exec.Command(cmdPath, args...)
		cmd.Dir = root
		cmd.Env = append(os.Environ(),
			"GH_PATH="+filepath.Join(root, "gh"),
			"GIT_PATH="+filepath.Join(root, "git"),
			// short circuits the git rev-parse the stubbed git can't answer
			"GIT_ROOT="+root,
		)

		ses, err := gexec.Start(cmd, GinkgoWriter, GinkgoWriter)
		Expect(err).NotTo(HaveOccurred())
		return ses
	}

	recorded := func() []string {
		data, err := os.ReadFile(calls)
		Expect(err).NotTo(HaveOccurred())
		return strings.Split(strings.TrimSpace(string(data)), "\n")
	}

	BeforeEach(func() {
		root = GinkgoT().TempDir()
		calls = filepath.Join(root, "calls")
		stub("gh", "", 0)
		stub("git", "", 0)
	})

	It("should list the subcommands", func() {
		ses := devctl("stack", "--help")

		Eventually(ses).Should(gexec.Exit(0))
		Expect(ses.Out).To(gbytes.Say("add"))
		Expect(ses.Out).To(gbytes.Say("edit"))
		Expect(ses.Out).To(gbytes.Say("save"))
		Expect(ses.Out).To(gbytes.Say("ship"))
	})

	It("should commit, rebase, and push", func() {
		ses := devctl("stack", "save", "-A", "-m", "Add the thing")

		Eventually(ses).Should(gexec.Exit(0))
		Expect(recorded()).To(Equal([]string{
			"git add -A",
			"git commit -m Add the thing",
			"gh stack rebase --upstack",
			"gh stack push",
		}))
	})

	It("should sync and submit", func() {
		ses := devctl("stack", "ship", "--open")

		Eventually(ses).Should(gexec.Exit(0))
		Expect(recorded()).To(Equal([]string{
			"gh stack sync",
			"gh stack submit --auto --open",
		}))
	})

	It("should exit with the code gh stack failed with", func() {
		stub("gh", "stack rebase --upstack", 3)

		ses := devctl("stack", "save")

		Eventually(ses).Should(gexec.Exit(3))
		Expect(recorded()).To(Equal([]string{"gh stack rebase --upstack"}))
	})
})
