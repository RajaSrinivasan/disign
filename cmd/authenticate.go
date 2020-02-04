package cmd

import (
	"fmt"
	"log"
	"path/filepath"

	"../impl"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
)

var public string

// authenticateCmd represents the authenticate command
var authenticateCmd = &cobra.Command{
	Use:   "authenticate",
	Short: "Verifies the signature file",
	Long: `
	
	Verifies that the file is authentic. 
	A digital signature is generated with the public part of the keypair and 
	compared to the contents of the signature file.
	
	For example:

	disign authenticate --public public-key-file filename
   `,
	Args: cobra.MinimumNArgs(1),
	Run:  authenticate,
}

func init() {
	rootCmd.AddCommand(authenticateCmd)
	h, err := homedir.Dir()
	if err != nil {
		log.Printf("%s\n", err)
	} else {
		public = filepath.Join(h, ".ssh", "id_rsa.pub")
	}
	authenticateCmd.PersistentFlags().StringVarP(&public, "public", "p", public, "Public key file name")
	authenticateCmd.MarkFlagRequired("public")

}

func authenticate(cmd *cobra.Command, args []string) {
	if Verbose {
		fmt.Printf("Public key %s\n", public)
		for _, arg := range args {
			fmt.Printf("Authenticate %s\n", arg)
		}
	}
	if len(public) == 0 {
		log.Fatal("Need to specify public key file\n")
	}
	impl.AuthenticateFiles(args, public)
}
