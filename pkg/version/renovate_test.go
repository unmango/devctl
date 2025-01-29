package version_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/spf13/afero"

	"github.com/unmango/devctl/pkg/version"
	"github.com/unmango/devctl/pkg/version/opts"
	"github.com/unmango/go/testing/gfs"
)

var _ = Describe("Renovate", func() {
	Describe("RenovateManager", func() {
		It("should set the manager type", func() {
			c := version.RenovateManager("test")

			Expect(c.CustomType).To(Equal("regex"))
		})

		It("should set the manager file match", func() {
			c := version.RenovateManager("test")

			Expect(c.FileMatch).To(ConsistOf(".versions/test"))
		})

		It("should set the manager match strings", func() {
			c := version.RenovateManager("test")

			Expect(c.MatchStrings).To(ConsistOf("(?<currentValue>+)"))
		})
	})

	Describe("RenovateConfig", func() {
		It("should add a custom manager", func() {
			c := version.RenovateConfig("test")

			Expect(c.CustomManagers).To(ConsistOf(
				version.RenovateManager("test"),
			))
		})
	})

	Describe("WriteRenovateConfig", func() {
		It("should write the given config", func() {
			c := version.RenovateConfig("test")
			fs := afero.NewMemMapFs()

			err := version.WriteRenovateConfig(c, opts.WithWriteFs(fs))

			Expect(err).NotTo(HaveOccurred())
			Expect(fs).To(gfs.ContainFile(version.RenovatePath))
		})
	})
})
