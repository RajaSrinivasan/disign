package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: version,
}

func init() {
	rootCmd.AddCommand(versionCmd)

}

func version(cmd *cobra.Command, args []string) {
	fmt.Printf("Version : %d.%d.%d\n", versionMajor, versionMinor, versionBuild)
	fmt.Printf("Built : %s\n", buildTime)
	fmt.Printf("Repo URL : %s\n", repoURL)
	fmt.Printf("Branch : %s\n", branchName)
	fmt.Printf("Commit Id : Short : %s Long %s\n", shortCommitId, longCommitId)
}
