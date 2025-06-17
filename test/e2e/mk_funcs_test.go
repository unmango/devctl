package e2e_test

import (
	"io/fs"
	"os"
	"path/filepath"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("MkFuncs", func() {
	var dir string

	BeforeEach(func() {
		fs, err := fs.Sub(testdata, "testdata")
		Expect(err).NotTo(HaveOccurred())

		dir = GinkgoT().TempDir()
		Expect(os.CopyFS(dir, fs)).To(Succeed())
	})

	It("should contain the file", func() {
		_, err := os.Open(filepath.Join(dir, "Makefile"))
		Expect(err).NotTo(HaveOccurred())
	})
})
