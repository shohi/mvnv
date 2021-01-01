package base

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/atrox/homedir"
	"github.com/shohi/mvnv/pkg/util"
)

const MvnMajorVersion = "3"
const RemoteURL = "https://archive.apache.org/dist/maven/maven-3/"

const MVNV_HOME = ".mvnv"
const MVN_BIN = "mvn"

const DIR_VERSIONS = "versions"
const DIR_DOWNLOADS = "downloads"
const DIR_TMP = "tmp"

func MvmHomeDir() string {
	home, _ := homedir.Dir()
	return filepath.Join(home, MVNV_HOME)
}

func VersionsDir() string {
	dir := filepath.Join(MvmHomeDir(), DIR_VERSIONS)
	_ = os.MkdirAll(dir, os.ModePerm)

	return dir
}

func DownloadsDir() string {
	dir := filepath.Join(MvmHomeDir(), DIR_DOWNLOADS)
	_ = os.MkdirAll(dir, os.ModePerm)

	return dir
}

func TmpDir() string {
	dir := filepath.Join(MvmHomeDir(), DIR_TMP)
	_ = os.MkdirAll(dir, os.ModePerm)

	return dir
}

func DefaultVersionFile() string {
	name := fmt.Sprintf(".%s-version", MVN_BIN)
	return filepath.Join(MvmHomeDir(), name)
}

// LoadInuseVersion loads the default activated version.
func LoadInuseVersion() string {
	v, err := util.LoadVersion(DefaultVersionFile())
	if err != nil {
		return ""
	}
	v = strings.Split(v, "\n")[0]

	return strings.TrimSpace(v)
}

func VersionBaseDir(version string) string {
	return filepath.Join(VersionsDir(), version)
}

func DownloadURLAndFileName(version string) (string, string) {
	filename := fmt.Sprintf("apache-maven-%s-bin.tar.gz", version)

	// https://archive.apache.org/dist/maven/maven-3/3.5.4/binaries/
	dlURL := fmt.Sprintf("%s%s/binaries/%s", RemoteURL, version, filename)

	return dlURL, filename
}

func DownloadBinaryFilePath(version string) string {
	filename := fmt.Sprintf("apache-maven-%s-bin.tar.gz", version)

	return filepath.Join(DownloadsDir(), filename)
}
