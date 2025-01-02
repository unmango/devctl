package version_test

import (
	"path/filepath"
	"testing/quick"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/unmango/devctl/pkg/version"
)

var _ = Describe("Version", func() {
	Describe("Regex", func() {
		DescribeTable("should match version-ish strings",
			Entry(nil, "v0.0.69"),
			Entry(nil, "0.0.69"),
			Entry(nil, "420.0.0"),
			func(input string) {
				matches := version.Regex.MatchString(input)

				Expect(matches).To(BeTrue())
			},
		)
	})

	Describe("RelPath", func() {
		It("should return the relative path to the given version file", func() {
			err := quick.Check(func(name string) bool {
				p := version.RelPath(name)
				return p == filepath.Join(version.DirName, name)
			}, nil)

			Expect(err).NotTo(HaveOccurred())
		})
	})

	Describe("Clean", func() {
		It("should trim leading 'v's", func() {
			v := version.Clean("v0.0.69")

			Expect(v).To(Equal("0.0.69"))
		})
	})
})
