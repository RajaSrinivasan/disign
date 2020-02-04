package cmd

import (
	"fmt"
	"log"
	"path/filepath"

	"github.com/RajaSrinivasan/disign/impl"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
)

var private string

// signCmd represents the sign command
var signCmd = &cobra.Command{
	Use:   "sign",
	Short: "Digitally sign the files",
	Long: `
	
	Digitally sign the file. 
	A digital signature is generated with the private part of the keypair and 
	is written to the signature file.
	
	For example:

	disign sign --private private-key-file filename
   `,
	Args: cobra.MinimumNArgs(1),
	Run:  sign,
}

func init() {
	rootCmd.AddCommand(signCmd)
	h, err := homedir.Dir()
	if err != nil {
		log.Printf("%s\n", err)
	} else {
		private = filepath.Join(h, ".ssh", "id_rsa")
	}
	signCmd.PersistentFlags().StringVarP(&private, "private", "p", private, "Private key file name")
	signCmd.MarkFlagRequired("private")
}

func sign(cmd *cobra.Command, args []string) {
	if Verbose {
		fmt.Printf("Private key file %s\n", private)
		for _, arg := range args {
			fmt.Printf("Sign %s\n", arg)
		}
	}
	if len(private) == 0 {
		log.Fatal("Need to specify private key file\n")
	}
	impl.Verbose = Verbose
	impl.SignFiles(args, private)

}
