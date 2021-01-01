package clean

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/shohi/mvnv/cmd/mvnv/base"
	"github.com/spf13/cobra"
)

func New() *cobra.Command {
	c := &cobra.Command{
		Use:   "clean",
		Short: "Remove downloaded source file",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				// clear all downloaded tars
				_ = clean("")
				return
			}

			// clear specified downloaded tar
			_ = clean(args[0])
		},
	}

	return c
}

func clean(version string) error {
	if version != "" {
		filename := base.DownloadBinaryFilePath(version)
		err := os.Remove(filename)
		if err == nil {
			fmt.Println("Remove", filepath.Base(filename))
		}

		return err
	}

	// clear all
	dlDir := base.DownloadsDir()
	entries, err := ioutil.ReadDir(dlDir)
	if err != nil {
		return err
	}

	for _, f := range entries {
		if err = os.Remove(filepath.Join(dlDir, f.Name())); err == nil {
			fmt.Println("Remove", f.Name())
		}
	}

	// last err
	return err

}
