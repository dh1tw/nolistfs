// NoListFileSystem is a custom filesystem implementation. It wraps
// around a "base" file system. It is mainly used with
// http.FileServer so that 404 status code is returned instead of
// a directory listing.
package nolistfs

import (
	"net/http"
	"os"
	"path/filepath"
)

type NoListFile struct {
	http.File
}

func (f NoListFile) Readdir(count int) ([]os.FileInfo, error) {
	return nil, nil
}

// NoListFileSystem is a custom filesystem implementation. It wraps
// around a "base" file system. It is mainly used with
// http.FileServer so that 404 status code is returned instead of
// a directory listing.
type NoListFileSystem struct {
	base http.FileSystem
}

func (nfs NoListFileSystem) Open(path string) (http.File, error) {
	f, err := nfs.base.Open(path)
	if err != nil {
		return nil, err
	}

	s, err := f.Stat()
	if err != nil {
		return nil, err
	}
	if s.IsDir() {
		index := filepath.Join(path, "index.html")
		if _, err := nfs.Open(index); err != nil {
			closeErr := f.Close()
			if closeErr != nil {
				return nil, closeErr
			}

			return nil, err
		}
	}
	return f, nil
}
