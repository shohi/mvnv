package util

import (
	"archive/tar"
	"errors"
	"os"
	"strings"

	"github.com/mholt/archiver/v3"
)

var (
	// used to stop tar file walker
	ErrStop       = errors.New("stop")
	ErrNotTarFile = errors.New("archiver: not tar file")
)

func GetTarFileRootDir(filename string) (string, error) {
	var rootDir string

	archiver.Walk(filename, func(f archiver.File) error {
		if h, ok := f.Header.(*tar.Header); ok {
			rootDir = h.Name
			return ErrStop
		}

		return nil
	})

	sep := string(os.PathSeparator)
	rootDir = strings.Split(rootDir, sep)[0]

	if rootDir == "" {
		return "", ErrNotTarFile
	}

	return rootDir, nil
}
