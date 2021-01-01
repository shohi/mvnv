package uninstall

import (
	"errors"
	"fmt"
	"os"

	"github.com/shohi/fileutil"
	"github.com/shohi/mvnv/cmd/mvnv/base"
	"github.com/spf13/cobra"
)

func New() *cobra.Command {
	c := &cobra.Command{
		Use:          "uninstall",
		Short:        "Uninstall a specific version",
		SilenceUsage: true,
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				fmt.Println("please specify version")
				return
			}

			// TODO: use os.Exit?
			err := uninstall(args[0])
			if err != nil {
				err = fmt.Errorf("uninstall %v failed: %v", args[0], err)
				fmt.Fprintln(os.Stderr, err)
				return
			}

			fmt.Printf("uninstall %v successfully\n", args[0])
		},
	}

	return c
}

func uninstall(version string) error {
	verDir := base.VersionBaseDir(version)

	if !fileutil.IsDir(verDir) {
		return errors.New("not installed")
	}

	return os.RemoveAll(verDir)
}
