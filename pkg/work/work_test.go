package work_test

import (
	"context"
	"os"
	"strings"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/unmango/devctl/pkg/work"
	"github.com/unmango/go/testing"
	"github.com/unmango/go/vcs/git"
)

var _ = Describe("Work", func() {
	var root string

	BeforeEach(func() {
		root = GinkgoT().TempDir()
	})

	Describe("Git", func() {
		It("should use the git path", func(ctx context.Context) {
			By("Initializing a git repo")
			testing.GitInit(ctx, root)

			c, err := work.Git(git.WithWorkingDirectory(ctx, root))

			Expect(err).NotTo(HaveOccurred())
			p := strings.TrimPrefix(c.Path(), "/private") // Mac crap
			Expect(p).To(Equal(root))
		})

		It("should error when the directory does not exist", func(ctx context.Context) {
			_, err := work.Git(git.WithWorkingDirectory(ctx, "blah"))

			Expect(err).To(MatchError("chdir blah: no such file or directory"))
		})
	})

	Describe("Cwd", func() {
		It("should use the current working directory", func() {
			wd, err := os.Getwd()
			Expect(err).NotTo(HaveOccurred())

			c, err := work.Cwd()

			Expect(err).NotTo(HaveOccurred())
			Expect(c.Path()).To(Equal(wd))
		})
	})

	Describe("Load", func() {
		It("should use the git path when in a repo", func(ctx context.Context) {
			By("Initializing a git repo")
			testing.GitInit(ctx, root)

			c, err := work.Load(git.WithWorkingDirectory(ctx, root))

			Expect(err).NotTo(HaveOccurred())
			p := strings.TrimPrefix(c.Path(), "/private") // Mac crap
			Expect(p).To(Equal(root))
		})

		It("should use the current working directory when not in a repo", func(ctx context.Context) {
			wd, err := os.Getwd()
			Expect(err).NotTo(HaveOccurred())

			c, err := work.Load(git.WithWorkingDirectory(ctx, root))

			Expect(err).NotTo(HaveOccurred())
			Expect(c.Path()).To(Equal(wd))
		})
	})
})
