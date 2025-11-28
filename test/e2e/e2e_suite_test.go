package e2e_test

import (
	"context"
	"embed"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"
	"github.com/unmango/go/vcs/git"
)

//go:embed testdata/happypath/*
//go:embed testdata/prefixed/*
//go:embed testdata/Makefile
var testdata embed.FS

var (
	cmdPath string
)

func TestE2e(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "E2e Suite")
}

var _ = BeforeSuite(func(ctx context.Context) {
	root, err := git.Root(ctx)
	Expect(err).NotTo(HaveOccurred())

	cmdPath, err = gexec.Build(root)
	Expect(err).NotTo(HaveOccurred())
})

var _ = AfterSuite(func() {
	gexec.CleanupBuildArtifacts()
})
