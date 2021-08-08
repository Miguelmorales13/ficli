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

// controllerCmd represents the controller command
var controllerCmd = &cobra.Command{
	Use:     "controller",
	Aliases: []string{"c"},
	Short:   "Creation a controller",
	Long:    `Creation a controller`,
	Args:    cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fileName := args[0]
		if isSingular := services.Plural.IsSingular(fileName); isSingular {
			fileName = services.Plural.Plural(fileName)
		}
		if flags, _ := cmd.Flags().GetBool("crud"); flags {
			services.TemplateControllerCrud(fileName)
		} else {
			services.TemplateControllerDefault(fileName)
		}
		cmd.Println("Done controller creation")
	},
}

func init() {
	controllerCmd.Flags().BoolP("crud", "c", false, "create a template to controller")
	generateCmd.AddCommand(controllerCmd)

}
