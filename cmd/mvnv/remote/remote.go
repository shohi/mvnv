package remote

import (
	"errors"

	"github.com/shohi/mvnv/cmd/mvnv/base"
	"github.com/shohi/mvnv/pkg/versions"
	"github.com/spf13/cobra"
)

func New() *cobra.Command {
	c := &cobra.Command{
		Use:          "list-remote",
		Short:        "List all installable versions",
		SilenceUsage: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			return listRemoteVersions()
		},
	}

	return c
}

func listRemoteVersions() error {
	c := versions.NewCollector(base.RemoteURL, base.MvnMajorVersion)
	vers, err := c.Remote()
	if err != nil {
		return err
	}

	if len(vers) == 0 {
		return errors.New("no version available")
	}

	versions.Print(vers, base.LoadInuseVersion())

	return nil
}
