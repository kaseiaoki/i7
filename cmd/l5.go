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
	"github.com/kaseiaoki/i7/listSegments"
	"github.com/spf13/cobra"
)

// l5Cmd represents the l5 command
var l5Cmd = &cobra.Command{
	Use:   "l5",
	Short: "ls",
	Long:  `ls command option`,
	Args: func(cmd *cobra.Command, args []string) error {
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			listSegments.Path(args[0])
			return
		}
		listSegments.Current()
		return
	},
}

func init() {
	rootCmd.AddCommand(l5Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// l5Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// l5Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
