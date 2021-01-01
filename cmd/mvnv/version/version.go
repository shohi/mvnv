package version

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
)

var version = "0.0.1"
var gitCommit string

type Info struct {
	Version   string `json:"version"`
	GitCommit string `json:"gitCommit"`
}

// New creates a new `version` subcommand.
func New() *cobra.Command {
	c := &cobra.Command{
		Use:   "version",
		Short: "Print version",
		RunE: func(cmd *cobra.Command, args []string) error {
			info := Info{
				Version:   version,
				GitCommit: gitCommit,
			}

			content, err := json.Marshal(info)
			if err != nil {
				return err
			}

			fmt.Println(string(content))
			return nil
		},
	}

	return c
}
