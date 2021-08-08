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

// handlerCmd represents the handler command
var handlerCmd = &cobra.Command{
	Use:   "handler",
	Short: "Generator of handler",
	Long:  `Generator of handler`,
	Run: func(cmd *cobra.Command, args []string) {
		if flags, _ := cmd.Flags().GetBool("error"); flags {
			services.TemplateHandlerError()
		} else {
			services.TemplateControllerDefault(args[0])
		}
		cmd.Println("Done controller creation")
	},
}

func init() {
	generateCmd.Flags().BoolP("error", "e", false, "create a template to error")

	generateCmd.AddCommand(handlerCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// handlerCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// handlerCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
