package use

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/shohi/mvnv/cmd/mvnv/base"
	"github.com/shohi/mvnv/pkg/versions"
	"github.com/spf13/cobra"
)

func New() *cobra.Command {
	c := &cobra.Command{
		Use:          "use",
		Short:        "Switch to specific version",
		SilenceUsage: true,
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				fmt.Println("please specify version")
				return
			}

			if err := use(args[0]); err != nil {
				fmt.Fprintln(os.Stderr, err)
				return
			}

			fmt.Printf("Done! Using %s version.\n", args[0])
		},
	}

	return c
}

func use(version string) error {
	version = strings.TrimSpace(version)
	vers, err := versions.Local(base.VersionsDir())
	if err != nil {
		return fmt.Errorf("failed to get installed version: %v", err)
	}

	if ok := versions.Find(vers, version); !ok {
		return fmt.Errorf("version %v does not exist, please install it first", version)
	}

	versionFile := base.DefaultVersionFile()
	err = ioutil.WriteFile(versionFile, []byte(version+"\n"), 0750)
	if err != nil {
		return fmt.Errorf("failed to set version: %v", err)
	}

	return nil
}
