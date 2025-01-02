package main_test

import (
	"context"
	"os"
	"os/exec"
	"path/filepath"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"
	"github.com/onsi/gomega/gexec"
	"github.com/spf13/afero"
	"github.com/unmango/devctl/pkg/version"
	"github.com/unmango/go/testing"
	"github.com/unmango/go/testing/gfs"
)

var _ = Describe("init", func() {
	var root string

	BeforeEach(func() {
		root = GinkgoT().TempDir()
	})

	Describe("version", func() {
		It("should initialize an inline version", func() {
			cmd := exec.Command(cmdPath, "init", "version", "blah", "0.0.69")
			cmd.Dir = root

			ses, err := gexec.Start(cmd, GinkgoWriter, GinkgoWriter)

			Expect(err).NotTo(HaveOccurred())
			Eventually(ses).Should(gexec.Exit(0))
			Expect(filepath.Join(root, version.DirName)).To(BeADirectory())
			Expect(afero.NewOsFs()).To(gfs.ContainFileWithBytes(
				filepath.Join(root, version.DirName, "blah"),
				[]byte("0.0.69"),
			))
		})

		It("should strip inline version prefixes", func() {
			cmd := exec.Command(cmdPath, "init", "version", "blah", "v0.0.69")
			cmd.Dir = root

			ses, err := gexec.Start(cmd, GinkgoWriter, GinkgoWriter)

			Expect(err).NotTo(HaveOccurred())
			Eventually(ses).Should(gexec.Exit(0))
			Expect(filepath.Join(root, version.DirName)).To(BeADirectory())
			Expect(afero.NewOsFs()).NotTo(gfs.ContainFileWithBytes(
				filepath.Join(root, version.DirName, "blah"),
				[]byte("v0.0.69"),
			))
		})

		It("should initialize an inline version with the auto source", func() {
			cmd := exec.Command(cmdPath, "init", "version", "blah", "v0.0.69", "--source", "auto")
			cmd.Dir = root

			ses, err := gexec.Start(cmd, GinkgoWriter, GinkgoWriter)

			Expect(err).NotTo(HaveOccurred())
			Eventually(ses).Should(gexec.Exit(0))
			Expect(filepath.Join(root, version.DirName)).To(BeADirectory())
			Expect(afero.NewOsFs()).To(gfs.ContainFileWithBytes(
				filepath.Join(root, version.DirName, "blah"),
				[]byte("0.0.69"),
			))
		})

		It("should error with an inline version and the github source", Pending, func() {
			cmd := exec.Command(cmdPath, "init", "version", "blah", "v0.0.69", "--source", "github")
			cmd.Dir = root

			ses, err := gexec.Start(cmd, GinkgoWriter, GinkgoWriter)

			Expect(err).NotTo(HaveOccurred())
			Eventually(ses).Should(gexec.Exit(1))
			Expect(ses.Err).To(gbytes.Say(`failed to get github version for: v0.0.69\n`))
		})

		It("should error when inline version does not have a name", func() {
			cmd := exec.Command(cmdPath, "init", "version", "v0.0.69")
			cmd.Dir = root

			ses, err := gexec.Start(cmd, GinkgoWriter, GinkgoWriter)

			Expect(err).NotTo(HaveOccurred())
			Eventually(ses).Should(gexec.Exit(1))
			Expect(ses.Err).To(gbytes.Say(`name is required\n`))
		})
	})

	It("should error when input is gibberish", func() {
		cmd := exec.Command(cmdPath, "init", "version", "blah-de-do-dah")
		cmd.Dir = root

		ses, err := gexec.Start(cmd, GinkgoWriter, GinkgoWriter)

		Expect(err).NotTo(HaveOccurred())
		Eventually(ses).Should(gexec.Exit(1))
		Expect(ses.Err).To(gbytes.Say(`unrecognized source: blah-de-do-dah\n`))
	})

	Context("Git repo", func() {
		BeforeEach(func(ctx context.Context) {
			By("Initializing a git repo in the working directory")
			testing.GitInit(ctx, root)
			Expect(os.Mkdir(filepath.Join(root, "subdir"), os.ModePerm)).To(Succeed())
		})

		It("should initialize an inline semver", func() {
			cmd := exec.Command(cmdPath, "init", "version", "blah", "v0.0.69")
			cmd.Dir = filepath.Join(root, "subdir")

			ses, err := gexec.Start(cmd, GinkgoWriter, GinkgoWriter)

			Expect(err).NotTo(HaveOccurred())
			Eventually(ses).Should(gexec.Exit(0))
			Expect(filepath.Join(root, version.DirName)).To(BeADirectory())
			Expect(afero.NewOsFs()).To(gfs.ContainFileWithBytes(
				filepath.Join(root, version.DirName, "blah"),
				[]byte("0.0.69"),
			))
		})
	})
})
