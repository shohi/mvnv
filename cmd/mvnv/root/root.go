package root

import (
	"os"

	"github.com/spf13/cobra"

	"github.com/shohi/mvnv/cmd/mvnv/base"
	"github.com/shohi/mvnv/cmd/mvnv/clean"
	"github.com/shohi/mvnv/cmd/mvnv/install"
	"github.com/shohi/mvnv/cmd/mvnv/list"
	"github.com/shohi/mvnv/cmd/mvnv/remote"
	"github.com/shohi/mvnv/cmd/mvnv/uninstall"
	"github.com/shohi/mvnv/cmd/mvnv/use"
	"github.com/shohi/mvnv/cmd/mvnv/version"
)

var logLevel string

// var conf = config.Config{}
var rootCmd = &cobra.Command{
	Use:   "mvnv",
	Short: "maven version manager",
}

// setupFlags sets flags for comand line
func setupFlags(cmd *cobra.Command) {
	cmd.PersistentFlags().String("loglevel", "INFO", "log level")
}

// Execute is the entrance.
func Execute() {
	// setup folders
	_ = os.MkdirAll(base.MvmHomeDir(), os.ModePerm)

	// init common flags
	setupFlags(rootCmd)

	// add subcommands
	rootCmd.AddCommand(list.New())
	rootCmd.AddCommand(remote.New())
	rootCmd.AddCommand(use.New())
	rootCmd.AddCommand(install.New())
	rootCmd.AddCommand(uninstall.New())
	rootCmd.AddCommand(clean.New())
	rootCmd.AddCommand(version.New())

	// run
	_ = rootCmd.Execute()
}
