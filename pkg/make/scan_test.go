package make_test

import (
	"bytes"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/unmango/devctl/pkg/make"
)

var _ = Describe("Scan", func() {
	It("should scan a target", func() {
		buf := bytes.NewBufferString(`target:`)
		s := make.NewScanner(buf)

		r := s.Scan()

		Expect(r).To(BeTrue())
		Expect(s.Node()).To(Equal(make.Rule{
			Target: []string{"target"},
		}))
	})

	It("should scan multiple targets", func() {
		buf := bytes.NewBufferString(`target:
target2:`)
		s := make.NewScanner(buf)

		r := s.Scan()

		Expect(r).To(BeTrue())
		Expect(s.Node()).To(Equal(make.Rule{
			Target: []string{"target", "target2"},
		}))
	})

	It("should ignore leading whitepace", func() {
		buf := bytes.NewBufferString(`
target:`)
		s := make.NewScanner(buf)

		r := s.Scan()

		Expect(r).To(BeTrue())
		Expect(s.Node()).To(Equal(make.Rule{
			Target: []string{"target", "target2"},
		}))
	})

	It("should split on newlines", func() {
		buf := bytes.NewBufferString("\n\n\n\n")
		s := make.NewScanner(buf)

		for s.Scan() {
		}

		Expect(s.Lines()).To(Equal([]string{
			"\n", "\n", "\n", "\n",
		}))
	})
})
