package work_test

import (
	"context"
	"os/exec"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/unmango/devctl/pkg/work"
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
			expected, err := gitInit(ctx)
			Expect(err).NotTo(HaveOccurred())
			GinkgoT().Chdir(expected)
			o := work.ChdirOptions{}

			p, err := o.Cwd(ctx)

			Expect(err).NotTo(HaveOccurred())
			Expect(p.Path()).To(Equal(expected))
		})
	})
})

func gitInit(ctx context.Context) (string, error) {
	cmd := exec.CommandContext(ctx, "git", "init")
	cmd.Dir = GinkgoT().TempDir()
	return cmd.Dir, cmd.Run()
}
