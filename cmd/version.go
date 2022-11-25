/*
Copyright © 2022 M.-Leander Reimer mario-leander.reimer@qaware.de

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

var version string
var commit string

// SetVersion set the application version for consumption in the output of the command.
func SetVersion(v string) {
	version = v
}

// SetCommit set the application commit for consumption in the output of the command.
func SetCommit(c string) {
	commit = c
}

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Display version information",
	Long:  "Display information on version number and Git commit",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("tf2eksctl %v %v\n", version, commit)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
