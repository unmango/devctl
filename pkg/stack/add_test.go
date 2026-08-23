package stack_test

import (
	"context"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/unmango/devctl/pkg/stack"
)

var _ = Describe("Add", func() {
	var (
		ctx     context.Context
		gh, git Stub
	)

	Context("in a stack", func() {
		BeforeEach(func() {
			ctx, gh, git = Stubs(nil, nil)
		})

		It("should add the branch", func() {
			err := stack.Add(ctx, "feat/api", stack.AddOptions{
				CommitOptions: stack.CommitOptions{
					All:     true,
					Message: "Add API routes",
				},
			})

			Expect(err).NotTo(HaveOccurred())
			Expect(git.Calls()).To(BeEmpty())
			Expect(gh.Calls()).To(Equal([]string{
				"stack|view|--json",
				"stack|add|-A|-m|Add API routes|feat/api",
			}))
		})

		It("should add the branch without committing", func() {
			err := stack.Add(ctx, "feat/api", stack.AddOptions{})

			Expect(err).NotTo(HaveOccurred())
			Expect(gh.Calls()).To(Equal([]string{
				"stack|view|--json",
				"stack|add|feat/api",
			}))
		})
	})

	Context("outside a stack", func() {
		BeforeEach(func() {
			ctx, gh, git = Stubs(map[string]Response{
				"stack view --json": {Code: stack.NotInStackCode},
			}, nil)
		})

		It("should initialize the stack and commit", func() {
			err := stack.Add(ctx, "feat/auth", stack.AddOptions{
				CommitOptions: stack.CommitOptions{
					All:     true,
					Message: "Add auth middleware",
				},
			})

			Expect(err).NotTo(HaveOccurred())
			Expect(gh.Calls()).To(Equal([]string{
				"stack|view|--json",
				"stack|init|feat/auth",
			}))
			Expect(git.Calls()).To(Equal([]string{
				"add|-A",
				"commit|-m|Add auth middleware",
			}))
		})

		It("should initialize the stack from the given base", func() {
			err := stack.Add(ctx, "feat/auth", stack.AddOptions{Base: "develop"})

			Expect(err).NotTo(HaveOccurred())
			Expect(gh.Calls()).To(Equal([]string{
				"stack|view|--json",
				"stack|init|--base|develop|feat/auth",
			}))
			Expect(git.Calls()).To(BeEmpty())
		})
	})

	It("should fail when the stack cannot be read", func() {
		ctx, gh, _ = Stubs(map[string]Response{
			"stack view --json": {Code: 4},
		}, nil)

		err := stack.Add(ctx, "feat/api", stack.AddOptions{})

		Expect(stack.ExitCode(err)).To(Equal(4))
		Expect(gh.Calls()).To(Equal([]string{"stack|view|--json"}))
	})
})
