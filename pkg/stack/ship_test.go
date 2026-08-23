package stack_test

import (
	"context"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/unmango/devctl/pkg/stack"
)

var _ = Describe("Ship", func() {
	var (
		ctx context.Context
		gh  Stub
	)

	BeforeEach(func() {
		ctx, gh, _ = Stubs(nil, nil)
	})

	It("should sync then submit", func() {
		err := stack.Ship(ctx, stack.ShipOptions{})

		Expect(err).NotTo(HaveOccurred())
		Expect(gh.Calls()).To(Equal([]string{
			"stack|sync",
			"stack|submit|--auto",
		}))
	})

	It("should prune and open", func() {
		err := stack.Ship(ctx, stack.ShipOptions{Open: true, Prune: true})

		Expect(err).NotTo(HaveOccurred())
		Expect(gh.Calls()).To(Equal([]string{
			"stack|sync|--prune",
			"stack|submit|--auto|--open",
		}))
	})

	It("should skip the sync", func() {
		err := stack.Ship(ctx, stack.ShipOptions{NoSync: true})

		Expect(err).NotTo(HaveOccurred())
		Expect(gh.Calls()).To(Equal([]string{"stack|submit|--auto"}))
	})

	It("should forward the remote", func() {
		err := stack.Ship(ctx, stack.ShipOptions{
			RemoteOptions: stack.RemoteOptions{Remote: "upstream"},
		})

		Expect(err).NotTo(HaveOccurred())
		Expect(gh.Calls()).To(Equal([]string{
			"stack|sync|--remote|upstream",
			"stack|submit|--auto|--remote|upstream",
		}))
	})

	It("should not submit when the sync fails", func() {
		ctx, gh, _ = Stubs(map[string]Response{
			"stack sync": {Code: 4},
		}, nil)

		err := stack.Ship(ctx, stack.ShipOptions{})

		Expect(stack.ExitCode(err)).To(Equal(4))
		Expect(gh.Calls()).To(Equal([]string{"stack|sync"}))
	})
})
