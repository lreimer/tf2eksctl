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
	"log"

	"github.com/lreimer/tf2ekscli/model"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

var clusterConfig model.ClusterConfig
var loggingTypes []string

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create eksctl YAML manifest for cluster",
	Long:  "Create eksctl YAML manifest for cluster from given arguments and Terraform state.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		clusterConfig.Metadata.Name = args[0]

		if len(loggingTypes) > 0 {
			clusterConfig.CloudWatch = model.CloudWatch{
				ClusterLogging: model.ClusterLogging{
					EnableTypes: loggingTypes,
				},
			}
		}

		d, err := yaml.Marshal(&clusterConfig)
		if err != nil {
			log.Fatalf("error: %v", err)
		}
		fmt.Printf("%s", string(d))
	},
}

func init() {
	rootCmd.AddCommand(createCmd)

	clusterConfig = model.NewClusterConfig()

	createCmd.Flags().StringVar(&clusterConfig.Metadata.Version, "version", "1.23", "EKS version to use, as semantic version")
	createCmd.Flags().StringVar(&clusterConfig.Metadata.Region, "region", "eu-central-1", "AWS region to use")

	createCmd.Flags().StringSliceVar(&loggingTypes, "enable-cluster-logging", nil, "Enable CloudWatch logging. Values: *, api, audit, authenticator, controllerManager, scheduler")

	createCmd.Flags().BoolVar(&clusterConfig.IAM.WithOIDC, "with-oidc", false, "Create OIDC provider for EKS cluster")
}
