package stack_test

import (
	"context"
	"errors"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/unmango/devctl/pkg/stack"
)

var _ = Describe("Stack", func() {
	Describe("CurrentBranch", func() {
		It("should return the branch name", func() {
			ctx, _, _ := Stubs(nil, map[string]Response{
				"rev-parse --abbrev-ref HEAD": {Out: "feat/ui"},
			})

			branch, err := stack.CurrentBranch(ctx)

			Expect(err).NotTo(HaveOccurred())
			Expect(branch).To(Equal("feat/ui"))
		})
	})

	Describe("InStack", func() {
		It("should be true when view succeeds", func() {
			ctx, _, _ := Stubs(nil, nil)

			Expect(stack.InStack(ctx)).To(BeTrue())
		})

		It("should be false when view exits with the not in a stack code", func() {
			ctx, _, _ := Stubs(map[string]Response{
				"stack view --json": {Code: stack.NotInStackCode},
			}, nil)

			Expect(stack.InStack(ctx)).To(BeFalse())
		})

		It("should fail on any other error", func() {
			ctx, _, _ := Stubs(map[string]Response{
				"stack view --json": {Code: 4},
			}, nil)

			_, err := stack.InStack(ctx)

			Expect(stack.ExitCode(err)).To(Equal(4))
		})
	})

	Describe("ExitCode", func() {
		It("should return -1 when the error is not from a process", func() {
			Expect(stack.ExitCode(errors.New("boom"))).To(Equal(-1))
		})
	})

	Describe("Gh", func() {
		It("should complain when gh is missing", func() {
			ctx := stack.WithGhPath(context.Background(), "")
			GinkgoT().Setenv("GH_PATH", "")
			GinkgoT().Setenv("PATH", GinkgoT().TempDir())

			err := stack.Gh(ctx, "stack", "view")

			Expect(err).To(MatchError(ContainSubstring("require the GitHub CLI")))
		})
	})
})
