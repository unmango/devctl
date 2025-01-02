package version_test

import (
	"testing/quick"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/unmango/go/devops/version"
)

var _ = Describe("Source", func() {
	Describe("GuessSource", func() {
		DescribeTable("should guess a semver version",
			Entry(nil, "v0.0.69"),
			Entry(nil, "0.0.69"),
			func(v string) {
				s, err := version.GuessSource(v)

				Expect(err).NotTo(HaveOccurred())
				Expect(s).To(Equal(version.String(v)))
			},
		)

		DescribeTable("should guess a github url",
			Entry(nil, "https://github.com/unmango/go"),
			Entry(nil, "github.com/unmango/go"),
			func(url string) {
				s, err := version.GuessSource(url)

				Expect(err).NotTo(HaveOccurred())
				Expect(s).To(Equal(version.GitHub(url)))
			},
		)

		It("should error on gibberish", func() {
			err := quick.Check(func(s string) bool {
				_, err := version.GuessSource(s)
				return err != nil
			}, nil)

			Expect(err).NotTo(HaveOccurred())
		})
	})

	Describe("String", func() {
		It("should use itself as the version", func() {
			err := quick.Check(func(input string) bool {
				s := version.String(input)
				v, err := s.Latest(nil)
				return err == nil && v == input
			}, nil)

			Expect(err).NotTo(HaveOccurred())
		})

		It("should not return a name", func() {
			err := quick.Check(func(s string) bool {
				v := version.String(s)
				_, err := v.Name(nil)
				return err != nil
			}, nil)

			Expect(err).NotTo(HaveOccurred())
		})
	})
})
