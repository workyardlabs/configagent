/*
Copyright © 2020 Workyard Labs <trevor@workyardlabs.com>

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
	"os"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/workyard/configagent/job"
)

var validateCmd = &cobra.Command{
	Use:   "validate",
	Short: "validate configuration file",
	Run: func(cmd *cobra.Command, args []string) {
		initLogging()

		file, err := os.Open(jobsFile)
		if err != nil {
			log.Fatal().Err(err)
		}

		jobs, err := job.ImportJobFile(file)
		if err != nil {
			log.Fatal().Err(err)
		}

		err = jobs.Validate()
		if err != nil {
			log.Fatal().Err(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(validateCmd)
}
