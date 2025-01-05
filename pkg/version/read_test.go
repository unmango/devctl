package version_test

import (
	"os"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/unmango/devctl/pkg/version"
)

var _ = Describe("Read", func() {
	Describe("Read", func() {
		It("should read happypath", func() {
			cwd, err := os.Getwd()
			Expect(err).NotTo(HaveOccurred())
			Expect(os.Chdir("testdata/happypath")).To(Succeed())
			DeferCleanup(os.Chdir, cwd)

			v, err := version.Read("test")

			Expect(err).NotTo(HaveOccurred())
			Expect(v).To(Equal("0.0.69"))
		})

		It("should error when dependency does not exist", func() {
			cwd, err := os.Getwd()
			Expect(err).NotTo(HaveOccurred())
			Expect(os.Chdir("testdata/happypath")).To(Succeed())
			DeferCleanup(os.Chdir, cwd)

			_, err = version.Read("wat")

			Expect(err).To(MatchError("dependency not found: .versions/wat"))
		})
	})
})
