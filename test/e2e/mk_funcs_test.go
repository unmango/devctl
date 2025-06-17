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

		err = os.Link(mkfuncsSo, filepath.Join(dir, "mk_funcs.so"))
		Expect(err).NotTo(HaveOccurred())
	})

	It("should contain the file", func() {
		dirs, err := os.ReadDir(dir)
		Expect(err).NotTo(HaveOccurred())
		Expect(dirs).To(BeEmpty())
	})
})
