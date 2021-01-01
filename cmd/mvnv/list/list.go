package list

import (
	"errors"

	"github.com/spf13/cobra"

	"github.com/shohi/mvnv/cmd/mvnv/base"
	"github.com/shohi/mvnv/pkg/versions"
)

func New() *cobra.Command {
	c := &cobra.Command{
		Use: "list",
		// Aliases:      []string{"ls"},
		Short:        "List all installed versions",
		SilenceUsage: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			return listVersions()
		},
	}

	return c
}

func listVersions() error {
	vers, err := versions.Local(base.VersionsDir())
	if err != nil {
		return err
	}

	if len(vers) == 0 {
		return errors.New("no maven installed yet")
	}

	versions.Print(vers, base.LoadInuseVersion())
	return nil
}
