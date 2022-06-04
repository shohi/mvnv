package version

import (
	"encoding/json"
	"fmt"
	"runtime/debug"

	"github.com/spf13/cobra"
)

var version = "0.0.2"
var gitCommit string

type Info struct {
	Version   string `json:"version"`
	GitCommit string `json:"gitCommit"`
	GoVersion string `json:"goVersion"`
}

// New creates a new `version` subcommand.
func New() *cobra.Command {
	c := &cobra.Command{
		Use:   "version",
		Short: "Print version",
		RunE: func(cmd *cobra.Command, args []string) error {
			build, ok := debug.ReadBuildInfo()
			if !ok {
				build = &debug.BuildInfo{}
			}

			info := Info{
				Version:   version,
				GitCommit: gitCommit,
				GoVersion: build.GoVersion,
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
