package e2e_test

import (
	"os"
	"os/exec"
	"path/filepath"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"
)

var _ = Describe("Install", func() {
	var root string

	BeforeEach(func() {
		root = GinkgoT().TempDir()
	})

	It("should install a tool from an archive", Label("E2E"), func() {
		cmd := exec.Command(cmdPath, "install", "devctl",
			"https://github.com/unmango/devctl/releases/download/v0.2.1/devctl_Linux_x86_64.tar.gz",
		)
		cmd.Dir = root

		ses, err := gexec.Start(cmd, GinkgoWriter, GinkgoWriter)

		Expect(err).NotTo(HaveOccurred())
		Eventually(ses, 10*time.Second).Should(gexec.Exit(0))

		binPath := filepath.Join(root, "bin", "devctl")
		Expect(binPath).To(BeARegularFile())

		cmd = exec.Command(binPath)
		Expect(cmd.Run()).To(Succeed())
	})

	It("should install a tool from a url", Label("E2E"), func() {
		cmd := exec.Command(cmdPath, "install", "mk_funcs.so",
			"https://github.com/unmango/devctl/releases/download/v0.2.1/mk_funcs.so",
		)
		cmd.Dir = root

		ses, err := gexec.Start(cmd, GinkgoWriter, GinkgoWriter)

		Expect(err).NotTo(HaveOccurred())
		Eventually(ses).Should(gexec.Exit(0))
		Expect(filepath.Join(root, "bin", "mk_funcs.so")).To(BeARegularFile())
	})

	When("a valid config file exists", func() {
		const configContent = `
tools:
  test:`

		var configPath string

		BeforeEach(func() {
			configPath = filepath.Join(root, "devctl.yml")
			err := os.WriteFile(configPath, []byte(configContent), os.ModePerm)
			Expect(err).NotTo(HaveOccurred())
		})

		It("should install the specified tools", Pending, func() {
			cmd := exec.Command(cmdPath, "install")
			cmd.Dir = root

			ses, err := gexec.Start(cmd, GinkgoWriter, GinkgoWriter)

			Expect(err).NotTo(HaveOccurred())
			Eventually(ses).Should(gexec.Exit(0))
			Expect(filepath.Join(root, "bin", "devctl")).To(BeARegularFile())
		})
	})
})
