package renovate_test

import (
	"encoding/json"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/unmango/devctl/pkg/renovate"
)

var _ = Describe("Schema", func() {
	Describe("CustomManager", func() {
		It("should marshal", func() {
			customType := renovate.ConfigCustomManagersElemCustomTypeRegex
			c := renovate.Config{CustomManagers: []renovate.ConfigCustomManagersElem{
				{CustomType: &customType},
			}}

			data, err := json.Marshal(c)

			Expect(err).NotTo(HaveOccurred())
			Expect(string(data)).To(Equal(`{"customManagers":[{"customType":"regex"}]}`))
		})
	})
})
