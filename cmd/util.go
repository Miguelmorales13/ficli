/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

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
	"github.com/Miguelmorales13/ficli/services"
	"github.com/spf13/cobra"
)

// utilCmd represents the util command
var utilCmd = &cobra.Command{
	Use:   "util",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if parse ,_:=cmd.Flags().GetBool("parseBody");parse {
			services.TemplateUtilsParseBody()
		} else if validations ,_:=cmd.Flags().GetBool("validations");validations {
			services.TemplateUtilsValidations()
		} else if jwt ,_:=cmd.Flags().GetBool("jwt");jwt {
			services.TemplateUtilsJWT()
		} else if emails ,_:=cmd.Flags().GetBool("emails");emails {
			services.TemplateUtilsEmails()
		} else {
			cmd.Println("Template not found")
		}
		
	},
}

func init() {
	utilCmd.Flags().BoolP("parseBody" ,"p",false,"Create a parser body for dtos")
	utilCmd.Flags().BoolP("validations" ,"v",false,"Create a validations for dtos")
	utilCmd.Flags().BoolP("jwt" ,"j",false,"Create a jwt util")
	utilCmd.Flags().BoolP("emails" ,"e",false,"Create a emails util")
	generateCmd.AddCommand(utilCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// utilCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// utilCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
