package version_test

import (
	"context"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/spf13/afero"

	"github.com/unmango/devctl/pkg/version"
	"github.com/unmango/go/testing/gfs"
)

type source struct {
	LatestFunc func(context.Context) (string, error)
	NameFunc   func(context.Context) (string, error)
}

func (s source) Latest(ctx context.Context) (string, error) {
	if s.LatestFunc != nil {
		return s.LatestFunc(ctx)
	} else {
		return "", nil
	}
}

func (s source) Name(ctx context.Context) (string, error) {
	if s.NameFunc != nil {
		return s.NameFunc(ctx)
	} else {
		return "", nil
	}
}

var _ = Describe("Init", func() {
	It("should create the .versions directory", func(ctx context.Context) {
		fs := afero.NewMemMapFs()
		src := source{}

		err := version.Init(ctx, "test", src,
			version.WithFs(fs),
		)

		Expect(err).NotTo(HaveOccurred())
		Expect(fs).To(gfs.ContainFile(".versions"))
	})

	It("should write the latest version", func(ctx context.Context) {
		fs := afero.NewMemMapFs()
		src := source{
			LatestFunc: func(ctx context.Context) (string, error) {
				return "0.0.69", nil
			},
		}

		err := version.Init(ctx, "test", src,
			version.WithFs(fs),
		)

		Expect(err).NotTo(HaveOccurred())
		Expect(fs).To(gfs.ContainFileWithBytes(
			".versions/test", []byte("0.0.69\n"),
		))
	})

	It("should use the source name when no name is provided", func(ctx context.Context) {
		fs := afero.NewMemMapFs()
		src := source{
			NameFunc: func(ctx context.Context) (string, error) {
				return "test-name", nil
			},
		}

		err := version.Init(ctx, "", src,
			version.WithFs(fs),
		)

		Expect(err).NotTo(HaveOccurred())
		Expect(fs).To(gfs.ContainFile(".versions/test-name"))
	})
})
