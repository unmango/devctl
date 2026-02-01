package renovate_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/unmango/devctl/pkg/renovate"
)

var _ = Describe("Renovate", func() {
	Describe("Unmarshal", func() {
		It("should work", func() {
			data := []byte(`{"branchName": "test"}`)

			c, err := renovate.Unmarshal(data)

			Expect(err).NotTo(HaveOccurred())
			Expect(c.BranchName).To(Equal("test"))
		})
	})
})
