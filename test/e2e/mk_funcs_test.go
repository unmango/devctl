package e2e_test

import (
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"
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
		cmd := exec.Command("make")
		cmd.Dir = dir

		ses, err := gexec.Start(cmd, GinkgoWriter, GinkgoWriter)

		Expect(err).NotTo(HaveOccurred())
		Eventually(ses).Should(gexec.Exit(0))
	})
})
