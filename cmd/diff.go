
// Copyright © 2019 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"errors"
	"log"
	"path/filepath"
	"github.com/spf13/cobra"
	"github.com/kaseiaoki/i7/fileDiff"
)

// worldCmd represents the world command
var diffCmd = &cobra.Command{
	Use:   "diff",
	Short: "this is plane file diff",
	Long:  `this is plane file diff`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 2 {
			return errors.New("Default arguments are two string value that is path for json file.")
		}
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		fa := args[0]
		fb := args[1]
	
		// file a, file b
		a, err := fileDiff.ReadFile(filepath.FromSlash(fa))
		if err != nil {
			log.Fatal(err)
		}
		b, err := fileDiff.ReadFile(filepath.FromSlash(fb))
		if err != nil {
			log.Fatal(err)
		}
	
		fileDiff.Diff(string(a), string(b))
		return nil
	},
}

func init() {
	rootCmd.AddCommand(diffCmd)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// worldCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// worldCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

