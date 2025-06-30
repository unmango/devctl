package tool

import (
	"archive/tar"
	"compress/gzip"
	"context"
	"fmt"
	"io"
	"net/http"
	"path"
	"strings"

	"github.com/charmbracelet/log"
	"github.com/unmango/devctl/pkg/work"
)

type Tool struct {
	Config

	Name string
}

func (t *Tool) Install(ctx context.Context, dir work.Directory) error {
	res, err := http.Get(t.Url)
	if err != nil {
		return err
	}

	var r io.Reader = res.Body
	defer res.Body.Close()

	if strings.HasSuffix(t.Url, ".tar.gz") {
		if r, err = Untar(t.Name, r); err != nil {
			return err
		}
	} else if path.Base(t.Url) != t.Name {
		return fmt.Errorf("%s not found at %s", t.Name, t.Url)
	}

	f, err := dir.Fs().Create(t.Name)
	if err != nil {
		return err
	}

	if _, err = io.Copy(f, r); err != nil {
		return err
	}

	if path.Ext(t.Name) == "" {
		if err = dir.Fs().Chmod(t.Name, 0o700); err != nil {
			return err
		}
	}

	return nil
}

func Untar(name string, r io.Reader) (io.Reader, error) {
	gzip, err := gzip.NewReader(r)
	if err != nil {
		return nil, err
	}

	tar := tar.NewReader(gzip)
	for {
		if h, err := tar.Next(); err != nil {
			break
		} else if h.Name == name {
			return tar, nil
		} else {
			log.Debugf("Ignoring %s", h.Name)
		}
	}

	return nil, fmt.Errorf("%s not found in archive", name)
}
