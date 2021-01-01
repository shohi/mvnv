package util

import (
	"errors"
	"io/ioutil"
	"path/filepath"
	"strings"

	"github.com/shohi/fileutil"
)

var (
	ErrFileNotExist  = errors.New("file: not exist")
	ErrFileNoContent = errors.New("file: no content")
)

// LoadVersion loads version info from file.
// Note only one version is allowed to be set in the file,
// and the first non-empty line is considered as version.
func LoadVersion(versionPath string) (string, error) {
	versionPath, _ = filepath.Abs(versionPath)

	if !fileutil.Exists(versionPath) {
		return "", ErrFileNotExist
	}

	version, err := ioutil.ReadFile(versionPath)
	if err != nil {
		return "", err
	}

	finalVersion := strings.Split(string(version), "\n")[0]
	finalVersion = strings.TrimSpace(finalVersion)
	if finalVersion == "" {
		return "", ErrFileNoContent
	}

	return finalVersion, nil
}

func IsNotExist(err error) bool {
	return err == ErrFileNotExist
}

func IsNoContent(err error) bool {
	return err == ErrFileNoContent
}
