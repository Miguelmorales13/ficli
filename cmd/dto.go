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

// dtoCmd represents the dto command
var dtoCmd = &cobra.Command{
	Use:   "dto",
	Short: "Creation of dto",
	Long:  `Creation of dto`,
	Run: func(cmd *cobra.Command, args []string) {
		fileName := args[0]
		if isSingular := services.Plural.IsSingular(fileName); isSingular {
			fileName = services.Plural.Plural(fileName)
		}
		if create, _ := cmd.Flags().GetBool("create"); create {
			services.TemplateCreateDto(fileName)
		} else if update, _ := cmd.Flags().GetBool("update"); update {
			services.TemplateUpdateDto(fileName)
		} else {
			services.TemplateDefaultDto(fileName)

		}
	},
}

func init() {
	dtoCmd.Flags().BoolP("create", "c", false, "Create a new dto with configuration the creation")
	dtoCmd.Flags().BoolP("update", "u", false, "Create a new dto with configuration the update")
	generateCmd.AddCommand(dtoCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// dtoCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// dtoCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
