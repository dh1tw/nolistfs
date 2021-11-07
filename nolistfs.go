// NoListFileSystem is a custom filesystem implementation. It
// follows the decorator pattern and wraps around a "base"
// file system. It is mainly used with http.FileServer so that
// 404 status code is returned instead of a directory listing.
package nolistfs

import (
	"net/http"
	"os"
	"path"
)

type NoListFile struct {
	http.File
}

// Readdir implements the Readdir function of http.FileSystem.
// In this specific case, it will always return nil, nil.
func (f NoListFile) Readdir(count int) ([]os.FileInfo, error) {
	return nil, nil
}

// New is a constructor method which returns the initialized
// No Listing File System.
func New(base http.FileSystem) http.FileSystem {
	nfs := &NoListFileSystem{
		base: base,
	}
	return nfs
}

// NoListFileSystem is a custom filesystem implementation. It wraps
// around a "base" file system. It is mainly used with
// http.FileServer so that 404 status code is returned instead of
// a directory listing.
type NoListFileSystem struct {
	base http.FileSystem
}

// Open implements Open method required by http.FileSystem. This
// function returns a 404 error when a trailing "/" is called.
func (nfs NoListFileSystem) Open(pathToOpen string) (http.File, error) {
	f, err := nfs.base.Open(pathToOpen)
	if err != nil {
		return nil, err
	}

	s, err := f.Stat()
	if err != nil {
		return nil, err
	}
	if s.IsDir() {
		index := path.Join(pathToOpen, "/index.html")
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
