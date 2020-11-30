package commands

import (
	"fmt"
	"runtime"

	"github.com/spf13/cobra"
)

var (
	// Automatically populated by goreleaser during build
	version = "v0.1.1"
	commit  = "?"
	date    = "?"
)

type versionRunner struct {
}

func NewVersionCommand() *cobra.Command {
	r := &versionRunner{}
	return &cobra.Command{
		Use:   "version",
		Short: "Print the current version",
		Run:   r.Run,
	}
}

func (r *versionRunner) Run(_ *cobra.Command, _ []string) {
	fmt.Printf("version=%s, commit=%s, buildDate=%s, os=%s, arch=%s\n", version, commit, date, runtime.GOOS, runtime.GOARCH)
}
