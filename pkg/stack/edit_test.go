package stack_test

import (
	"context"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/unmango/devctl/pkg/stack"
)

var _ = Describe("Edit", func() {
	var (
		ctx     context.Context
		gh, git Stub
	)

	BeforeEach(func() {
		ctx, gh, git = Stubs(nil, map[string]Response{
			"rev-parse --abbrev-ref HEAD": {Out: "feat/ui"},
		})
	})

	It("should commit to the branch and return to the current one", func() {
		err := stack.Edit(ctx, "feat/auth", stack.EditOptions{
			CommitOptions: stack.CommitOptions{
				All:     true,
				Message: "Fix the middleware",
			},
		})

		Expect(err).NotTo(HaveOccurred())
		Expect(gh.Calls()).To(Equal([]string{
			"stack|checkout|feat/auth",
			"stack|rebase|--upstack",
			"stack|checkout|feat/ui",
			"stack|push",
		}))
		Expect(git.Calls()).To(Equal([]string{
			"rev-parse|--abbrev-ref|HEAD",
			"add|-A",
			"commit|-m|Fix the middleware",
		}))
	})

	It("should not check out the branch it is already on", func() {
		err := stack.Edit(ctx, "feat/ui", stack.EditOptions{})

		Expect(err).NotTo(HaveOccurred())
		Expect(gh.Calls()).To(Equal([]string{
			"stack|rebase|--upstack",
			"stack|push",
		}))
	})

	It("should skip pushing", func() {
		err := stack.Edit(ctx, "feat/auth", stack.EditOptions{NoPush: true})

		Expect(err).NotTo(HaveOccurred())
		Expect(gh.Calls()).To(Equal([]string{
			"stack|checkout|feat/auth",
			"stack|rebase|--upstack",
			"stack|checkout|feat/ui",
		}))
	})

	It("should forward the remote", func() {
		err := stack.Edit(ctx, "feat/auth", stack.EditOptions{
			RemoteOptions: stack.RemoteOptions{Remote: "upstream"},
		})

		Expect(err).NotTo(HaveOccurred())
		Expect(gh.Calls()).To(ContainElements(
			"stack|rebase|--upstack|--remote|upstream",
			"stack|push|--remote|upstream",
		))
	})

	It("should stay put when the checkout fails", func() {
		ctx, gh, _ = Stubs(map[string]Response{
			"stack checkout feat/auth": {Code: 1},
		}, map[string]Response{
			"rev-parse --abbrev-ref HEAD": {Out: "feat/ui"},
		})

		err := stack.Edit(ctx, "feat/auth", stack.EditOptions{})

		Expect(stack.ExitCode(err)).To(Equal(1))
		Expect(gh.Calls()).To(Equal([]string{"stack|checkout|feat/auth"}))
	})
})
