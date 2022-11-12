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
	"github.com/spf13/cobra"
)

var banner string = `
                                            
			      AW            
'7MM"""YMM  db '7MMF'   '7MF',M''7MM"""YMM  
  MM    '7 ;MM:  'MA     ,V  MV   MM    '7  
  MM   d  ,V^MM.  VM:   ,V  AW    MM   d    
  MM""MM ,M  'MM   MM.  M' ,M'    MMmmMM    
  MM   Y AbmmmqMA  'MM A'  MV     MM   Y  , 
  MM    A'     VML  :MM;  AW      MM     ,M 
.JMML..AMA.   .AMMA. VF  ,M'    .JMMmmmmMMM 
                         MV                 
                        AW                  

`

var rootCmd = &cobra.Command{
	Use:   "fave",
	Short: "Find a vulnerability/enumeration",
	Long: banner + `
FAV/E is a CLI tool designed to enable quicker searches of the NIST CVE database.
This tool will give a real-time look at publicly known vulnerabilities for a given
keyword or phrase, all from the command line. 

At this time, internet access is required for the tool to gather data. An offline 
version may be released at a later date, but tools like searchsploit ultimately 
perform the same function that an offline version of this would solve.
`,
	Run: func(cmd *cobra.Command, args []string) {
	},
	Version: "1.1.0",
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	rootCmd.AddCommand(searchCmd)
	rootCmd.AddCommand(describeCmd)
	rootCmd.CompletionOptions.DisableDefaultCmd = true
}
