/*
Copyright © 2019 R Srinivasan rs@toprllc.com

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"

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
	signCmd.PersistentFlags().StringVarP(&private, "private", "p", "", "Private key file name")
	signCmd.MarkFlagRequired("private")
}

func sign(cmd *cobra.Command, args []string) {
	if Verbose {
		fmt.Printf("Public key %s\n", public)
		for _, arg := range args {
			fmt.Printf("Sign %s\n", arg)
		}
	}
}
