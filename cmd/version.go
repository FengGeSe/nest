package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	_version   = "unknown"
	_gitCommit = "unknown"
	_goVersion = "unknown"
	_buildTime = "unknown"
	_osArch    = "unknown"

	version = ""
)

func init() {
	rootCmd.AddCommand(versionCmd)

	version = fmt.Sprintf(`   Version:          %s
   Go version:       %s
   Git commit:       %s
   Built:            %s
   OS/Arch:          %s`, _version, _goVersion, _gitCommit, _buildTime, _osArch)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show Nest version info",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Println(version)
	},
}
