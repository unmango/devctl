package main_test

import (
	"embed"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"
)

//go:embed testdata/happypath/*
//go:embed testdata/prefixed/*
var testdata embed.FS

var cmdPath string

func TestDevctl(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Devctl Suite")
}

var _ = BeforeSuite(func() {
	p, err := gexec.Build("main.go")
	Expect(err).NotTo(HaveOccurred())
	cmdPath = p
})

var _ = AfterSuite(func() {
	gexec.CleanupBuildArtifacts()
})
