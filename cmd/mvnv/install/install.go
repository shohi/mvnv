package install

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/mholt/archiver/v3"
	"github.com/shohi/fileutil"
	"github.com/spf13/cobra"

	"github.com/shohi/mvnv/cmd/mvnv/base"
	"github.com/shohi/mvnv/pkg/check"
	"github.com/shohi/mvnv/pkg/download"
	"github.com/shohi/mvnv/pkg/util"
	"github.com/shohi/mvnv/pkg/versions"
)

func New() *cobra.Command {
	c := &cobra.Command{
		Use:   "install",
		Short: "Install a specific version",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				fmt.Println("please specify version")
				return
			}

			err := install(args[0])
			if err != nil {
				err = fmt.Errorf("install %v failed: %v", args[0], err)
				fmt.Fprintln(os.Stderr, err)
				return
			}

			fmt.Printf("install %v successfully\n", args[0])
		},
	}

	return c
}

func install(version string) error {
	var versionsDir = base.VersionsDir()

	if err := validateVersion(version); err != nil {
		return err
	}

	// use checksum to avoid duplicated download
	dlURL, _ := base.DownloadURLAndFileName(version)
	targetPath := base.DownloadBinaryFilePath(version)

	if ok := fileutil.Exists(targetPath); ok {
		matched, _ := check.VerifyChecksum(targetPath,
			dlURL, base.ChecksumAlgorithm(version))
		if matched {
			goto unzip
		} else {
			// delete corrupted file
			_ = os.Remove(targetPath)
		}

	}

	if _, err := download.Download(dlURL, targetPath); err != nil {
		return err
	}

unzip:
	if err := archiver.Unarchive(targetPath, versionsDir); err != nil {
		return err
	}

	verDir := base.VersionBaseDir(version)
	if err := os.MkdirAll(verDir, os.ModePerm); err != nil {
		return err
	}

	// remove all content under folder
	_ = os.RemoveAll(verDir)

	tarRootDir, err := util.GetTarFileRootDir(targetPath)
	if err != nil {
		return err
	}

	if err = os.Rename(filepath.Join(versionsDir, tarRootDir), verDir); err != nil {
		return err
	}

	return nil
}

func validateVersion(ver string) error {
	c := versions.NewCollector(base.RemoteURL, base.MvnMajorVersion)
	vers, err := c.Remote()
	if err != nil {
		return err
	}
	if ok := versions.Find(vers, ver); !ok {
		return errors.New("invalid")
	}

	return nil
}
