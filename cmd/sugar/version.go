package sugar

import (
	"fmt"

	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "print the version of Gins",
	Long:  "print the version of Gins",
	Run:   DisplayVersion,
}

func DisplayVersion(cmd *cobra.Command, args []string) {
	fmt.Printf("APPVersion: %v  BuildTime: %v  GitCommit: %v\n",
		APPVersion,
		BuildTime,
		GitCommit)
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
