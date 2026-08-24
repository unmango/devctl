package stack_test

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/unmango/devctl/pkg/stack"
)

func TestStack(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Stack Suite")
}

// Stubs returns a context resolving gh and git to newly created stubs
func Stubs(gh, git map[string]Response) (context.Context, Stub, Stub) {
	ghStub, gitStub := NewStub("gh", gh), NewStub("git", git)
	ctx := stack.WithGhPath(
		stack.WithGitPath(context.Background(), gitStub.Path),
		ghStub.Path,
	)

	return ctx, ghStub, gitStub
}

// A Response is what a stub prints and exits with for a given invocation
type Response struct {
	Out  string
	Code int
}

// A Stub is a fake gh or git recording every invocation it receives
type Stub struct {
	Path string
	log  string
}

// NewStub writes an executable recording its arguments to a log file.
//
// Keys of responses are arguments joined by a space, matched exactly against
// the invocation. Anything unmatched prints nothing and succeeds.
func NewStub(name string, responses map[string]Response) Stub {
	dir := GinkgoT().TempDir()
	stub := Stub{
		Path: filepath.Join(dir, name),
		log:  filepath.Join(dir, name+".log"),
	}

	var cases strings.Builder
	for args, res := range responses {
		if res.Out != "" {
			fmt.Fprintf(&cases, "%q)\n\tprintf '%%s\\n' %q\n\texit %d\n\t;;\n",
				args, res.Out, res.Code,
			)
		} else {
			fmt.Fprintf(&cases, "%q)\n\texit %d\n\t;;\n", args, res.Code)
		}
	}

	// arguments are logged pipe separated so that a commit message containing
	// spaces stays a single field
	script := fmt.Sprintf(`#!/bin/sh
line=""
sep=""
for arg in "$@"; do
	line="$line$sep$arg"
	sep="|"
done
printf '%%s\n' "$line" >> %q
case "$*" in
%sesac
exit 0
`, stub.log, cases.String())

	Expect(os.WriteFile(stub.Path, []byte(script), 0o755)).To(Succeed())
	return stub
}

// Calls returns every invocation the stub received, arguments pipe separated
func (s Stub) Calls() []string {
	data, err := os.ReadFile(s.log)
	if os.IsNotExist(err) {
		return nil
	}

	Expect(err).NotTo(HaveOccurred())
	return strings.Split(strings.TrimSpace(string(data)), "\n")
}
