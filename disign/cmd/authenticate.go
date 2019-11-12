/*
Copyright © 2019 NAME HERE <EMAIL ADDRESS>

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
	"log"

	"../impl"
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

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// authenticateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	authenticateCmd.PersistentFlags().StringVarP(&public, "public", "p", "", "Public key file name")
	authenticateCmd.MarkFlagRequired("public")

	// is called directly, e.g.:
	// authenticateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
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
	impl.Authenticate(public, args)
}
