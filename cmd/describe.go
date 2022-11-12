/*
Copyright © 2022 Tony West

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
	"github.com/un4gi/fave/api"
)

func init() {}

var describeCmd = &cobra.Command{
	Use:   "describe <CVE-ID>",
	Short: "Display detailed information about a specific CVE ID.",
	Long: banner + `
Once an interesting CVE is found, you can pass the CVE-ID to the describe command to gather a more detailed description with references.
This should be in the format <CVE-XXXX-XXXXX>. An example is as follows:

fave describe CVE-1900-01234
`,
	Args: cobra.MatchAll(cobra.MinimumNArgs(1), cobra.MaximumNArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(banner)
		api.DescribeCVE(args[0])
	},
}
