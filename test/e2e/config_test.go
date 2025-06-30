package e2e_test

import (
	"os"
	"os/exec"
	"path/filepath"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"
)

var _ = Describe("Config", func() {
	var root string

	BeforeEach(func() {
		root = GinkgoT().TempDir()
	})

	When("a valid config file exists", func() {
		var configPath string

		BeforeEach(func() {
			configPath = filepath.Join(root, "devctl.yml")
			err := os.WriteFile(configPath, []byte{}, os.ModePerm)
			Expect(err).NotTo(HaveOccurred())
		})

		It("should verify", func() {
			cmd := exec.Command(cmdPath, "config", "verify")
			cmd.Dir = root

			ses, err := gexec.Start(cmd, GinkgoWriter, GinkgoWriter)

			Expect(err).NotTo(HaveOccurred())
			Eventually(ses).Should(gexec.Exit(0))
		})
	})
})
