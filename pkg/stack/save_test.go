package stack_test

import (
	"context"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/unmango/devctl/pkg/stack"
)

var _ = Describe("Save", func() {
	var (
		ctx     context.Context
		gh, git Stub
	)

	BeforeEach(func() {
		ctx, gh, git = Stubs(nil, nil)
	})

	It("should rebase and push", func() {
		err := stack.Save(ctx, stack.SaveOptions{})

		Expect(err).NotTo(HaveOccurred())
		Expect(git.Calls()).To(BeEmpty())
		Expect(gh.Calls()).To(Equal([]string{
			"stack|rebase|--upstack",
			"stack|push",
		}))
	})

	It("should stage and commit before rebasing", func() {
		err := stack.Save(ctx, stack.SaveOptions{
			CommitOptions: stack.CommitOptions{
				All:     true,
				Message: "Add the thing",
			},
		})

		Expect(err).NotTo(HaveOccurred())
		Expect(git.Calls()).To(Equal([]string{
			"add|-A",
			"commit|-m|Add the thing",
		}))
		Expect(gh.Calls()).To(Equal([]string{
			"stack|rebase|--upstack",
			"stack|push",
		}))
	})

	It("should stage tracked files only", func() {
		err := stack.Save(ctx, stack.SaveOptions{
			CommitOptions: stack.CommitOptions{Update: true},
		})

		Expect(err).NotTo(HaveOccurred())
		Expect(git.Calls()).To(Equal([]string{"add|-u"}))
	})

	It("should skip pushing", func() {
		err := stack.Save(ctx, stack.SaveOptions{NoPush: true})

		Expect(err).NotTo(HaveOccurred())
		Expect(gh.Calls()).To(Equal([]string{"stack|rebase|--upstack"}))
	})

	It("should forward the remote", func() {
		err := stack.Save(ctx, stack.SaveOptions{
			RemoteOptions: stack.RemoteOptions{Remote: "upstream"},
		})

		Expect(err).NotTo(HaveOccurred())
		Expect(gh.Calls()).To(Equal([]string{
			"stack|rebase|--upstack|--remote|upstream",
			"stack|push|--remote|upstream",
		}))
	})

	It("should stop when the rebase fails", func() {
		ctx, gh, _ = Stubs(map[string]Response{
			"stack rebase --upstack": {Code: 3},
		}, nil)

		err := stack.Save(ctx, stack.SaveOptions{})

		Expect(stack.ExitCode(err)).To(Equal(3))
		Expect(gh.Calls()).To(Equal([]string{"stack|rebase|--upstack"}))
	})
})
