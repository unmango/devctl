package work_test

import (
	"context"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/unmango/devctl/pkg/work"
	"github.com/unmango/go/vcs/git"
)

var _ = Describe("Cmd", func() {
	Describe("ChdirOptions", func() {
		It("should return the chdir when it is defined", func(ctx context.Context) {
			o := work.NewChdirOptions("blah")

			p, err := o.Cwd(ctx)

			Expect(err).NotTo(HaveOccurred())
			Expect(p.Path()).To(Equal("blah"))
		})

		It("should return the git path with chdir is empty", func(ctx context.Context) {
			expected, err := git.Root(ctx)
			Expect(err).NotTo(HaveOccurred())
			o := work.ChdirOptions{}

			p, err := o.Cwd(ctx)

			Expect(err).NotTo(HaveOccurred())
			Expect(p.Path()).To(Equal(expected))
		})
	})
})
