package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var cfgFile string
var Verbose bool

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "disign",
	Short: "Digitally sign a file or verify the signature",
	Long: `Generate a cyrptographic signature of a file.
  
  The signature is generated using the private part of the keypair. 
  
  For verification, the public part of the keypair is used. 
  
  File naming convention:

  The signature file is named <basename>.sig 

  disign -private file1, ...
  disign -authenticate file1, file2, ...

  The signature files and the base files are in the same folder
  `,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "verbose output")
}
