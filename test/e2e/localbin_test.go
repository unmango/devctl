package e2e_test

import (
	"os/exec"
	"path/filepath"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"
	"github.com/onsi/gomega/gexec"
)

var _ = Describe("localbin", func() {
	var root string

	BeforeEach(func() {
		root = GinkgoT().TempDir()
	})

	It("should print the path", func() {
		cmd := exec.Command(cmdPath, "localbin")
		cmd.Dir = root

		ses, err := gexec.Start(cmd, GinkgoWriter, GinkgoWriter)

		Expect(err).NotTo(HaveOccurred())
		Eventually(ses.Out).Should(gbytes.Say(filepath.Join(root, "bin")))
	})
})
