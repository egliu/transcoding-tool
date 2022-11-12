package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var (
	version    = "0.1"
	versionCmd = &cobra.Command{
		Use:   "version",
		Short: "Print the version number of transcoding-tool",
		Long:  `All software has versions. This is transcoding-tool's`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(version)
		},
	}
)

func init() {
	rootCmd.AddCommand(versionCmd)
}
