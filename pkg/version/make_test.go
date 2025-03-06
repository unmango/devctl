package version_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/spf13/afero"

	"github.com/unmango/aferox/testing/gfs"
	"github.com/unmango/devctl/pkg/version"
	"github.com/unmango/devctl/pkg/version/opts"
)

var _ = Describe("Make", func() {
	It("should return the correct makefile variable", func() {
		text := version.MakefileVar("test")

		Expect(text).To(Equal(
			"TEST_VERSION := $(shell devctl version test)\n",
		))
	})

	It("should write the makefile", func() {
		fs := afero.NewMemMapFs()

		err := version.WriteMakefile("test",
			opts.WithWriteFs(fs),
		)

		Expect(err).ToNot(HaveOccurred())
		Expect(fs).To(gfs.ContainFile(".versions/test.mk"))
	})
})
