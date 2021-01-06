package wrapper

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"

	"github.com/shohi/mvnv/pkg/util"
)

// TODO: remove hardcoded snippets
// get the requested binary version from version file
// or use default version
func resolveVersion(baseDir, binName string) (string, error) {
	// get local version if set
	localVersionPath := fmt.Sprintf(".%s_version", binName)
	localVersion, err := util.LoadVersion(localVersionPath)
	if err != nil && !util.IsNotExist(err) {
		return "", err
	}

	if localVersion != "" {
		return localVersion, nil
	}

	// check default version if local one is not set
	defaultVersionPath := fmt.Sprintf("%s/.%s-version", baseDir, binName)

	defaultVersion, err := util.LoadVersion(defaultVersionPath)
	if err != nil {
		err := fmt.Errorf("please set %v version first", binName)
		return "", err
	}

	return defaultVersion, nil
}

func Wrapper(baseDir, binName string) {
	version, err := resolveVersion(baseDir, binName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		return
	}

	var fileExt string
	if runtime.GOOS == "windows" {
		fileExt = ".exe"
	}

	// TODO: refactor
	// bin path - ~/<baseDir>/versions/<version>/<binName><fileExt>
	bin := fmt.Sprintf("%s/versions/%s/bin/%s", baseDir, version, binName)
	bin, _ = filepath.Abs(bin)
	bin += fileExt

	cmd := exec.Command(bin, os.Args[1:]...)
	cmd.Env = os.Environ()
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr

	err = cmd.Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
	}
}
