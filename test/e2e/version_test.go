package e2e_test

import (
	"context"
	"os"
	"os/exec"
	"path/filepath"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"
	"github.com/onsi/gomega/gexec"
	"github.com/unmango/go/testing"
)

var _ = Describe("version", func() {
	var root string

	BeforeEach(func() {
		By("Creating a working directory")
		root = GinkgoT().TempDir()

		By("Copying testdata")
		Expect(os.CopyFS(root, testdata)).To(Succeed())
	})

	It("should magically print the version when given a version path", func() {
		cmd := exec.Command(cmdPath, ".versions/test")
		cmd.Dir = filepath.Join(root, "testdata", "happypath")

		ses, err := gexec.Start(cmd, GinkgoWriter, GinkgoWriter)

		Expect(err).NotTo(HaveOccurred())
		Eventually(ses).Should(gexec.Exit(0))
		Expect(ses.Out).To(gbytes.Say("v0.0.69"))
	})

	It("should print the version", func() {
		cmd := exec.Command(cmdPath, "version", "test")
		cmd.Dir = filepath.Join(root, "testdata", "happypath")

		ses, err := gexec.Start(cmd, GinkgoWriter, GinkgoWriter)

		Expect(err).NotTo(HaveOccurred())
		Eventually(ses).Should(gexec.Exit(0))
		Expect(ses.Out).To(gbytes.Say("^0.0.69\n"))
	})

	DescribeTable("should change to the specified directory",
		Entry("Long option", "--chdir"),
		Entry("Short option", "-C"),
		func(opt string) {
			cmd := exec.Command(cmdPath, "version", "test",
				opt, filepath.Join(root, "testdata", "happypath"),
			)

			ses, err := gexec.Start(cmd, GinkgoWriter, GinkgoWriter)

			Expect(err).NotTo(HaveOccurred())
			Eventually(ses).Should(gexec.Exit(0))
			Expect(ses.Out).To(gbytes.Say("^0.0.69\n"))
		},
	)

	It("should clean a prefixed version", func() {
		cmd := exec.Command(cmdPath, "version", "test")
		cmd.Dir = filepath.Join(root, "testdata", "prefixed")

		ses, err := gexec.Start(cmd, GinkgoWriter, GinkgoWriter)

		Expect(err).NotTo(HaveOccurred())
		Eventually(ses).Should(gexec.Exit(0))
		Expect(ses.Out).To(gbytes.Say("^0.0.69\n"))
	})

	It("should print the prefixed version", func() {
		cmd := exec.Command(cmdPath, "version", "test", "--prefixed")
		cmd.Dir = filepath.Join(root, "testdata", "happypath")

		ses, err := gexec.Start(cmd, GinkgoWriter, GinkgoWriter)

		Expect(err).NotTo(HaveOccurred())
		Eventually(ses).Should(gexec.Exit(0))
		Expect(ses.Out).To(gbytes.Say("0.0.69\n"))
	})

	It("should should error when dependency does not exist", func() {
		cmd := exec.Command(cmdPath, "version", "wat")
		cmd.Dir = filepath.Join(root, "testdata", "happypath")

		ses, err := gexec.Start(cmd, GinkgoWriter, GinkgoWriter)

		Expect(err).NotTo(HaveOccurred())
		Eventually(ses).Should(gexec.Exit(1))
		Expect(ses.Err).To(gbytes.Say("dependency not found: .versions/wat\n"))
	})

	Context("Git", func() {
		BeforeEach(func(ctx context.Context) {
			By("Initializing a git repo in the working directory")
			testing.GitInit(ctx, root)
			err := os.Mkdir(filepath.Join(root, "subdir"), os.ModePerm)
			Expect(err).NotTo(HaveOccurred())

			By("Creating a test version file")
			err = os.Mkdir(filepath.Join(root, ".versions"), os.ModePerm)
			Expect(err).NotTo(HaveOccurred())
			err = os.WriteFile(filepath.Join(root, ".versions", "test"), []byte("0.0.69"), os.ModePerm)
			Expect(err).NotTo(HaveOccurred())
		})

		It("should print the version", func() {
			cmd := exec.Command(cmdPath, "version", "test")
			cmd.Dir = root

			ses, err := gexec.Start(cmd, GinkgoWriter, GinkgoWriter)

			Expect(err).NotTo(HaveOccurred())
			Eventually(ses).Should(gexec.Exit(0))
			Expect(ses.Out).To(gbytes.Say("^0.0.69\n"))
		})

		It("should print the version from a subdir", func() {
			cmd := exec.Command(cmdPath, "version", "test")
			cmd.Dir = filepath.Join(root, "subdir")

			ses, err := gexec.Start(cmd, GinkgoWriter, GinkgoWriter)

			Expect(err).NotTo(HaveOccurred())
			Eventually(ses).Should(gexec.Exit(0))
			Expect(ses.Out).To(gbytes.Say("^0.0.69\n"))
		})
	})
})
