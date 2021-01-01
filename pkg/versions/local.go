package versions

import "github.com/shohi/fileutil"

// Local get all installed versions which locate at baseDir.
// e.g. <baseDir>/{3.6.3, 3.5.4}
func Local(baseDir string) ([]string, error) {
	return fileutil.AllSubfolders(baseDir)
}
