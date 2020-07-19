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
	"time"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/workyard/configagent/job"
)

var (
	tickSeconds int
	daemon      bool
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "run configuration check and replacement",
	Run: func(cmd *cobra.Command, args []string) {
		initLogging()

		jobs, err := job.LoadJobs(jobsFile)
		if err != nil {
			log.Fatal().Err(err)
		}

		if daemon {
			log.Info().Msgf("daemon started: tick(%d)", tickSeconds)

			ticker := time.NewTicker(time.Duration(tickSeconds) * time.Second)

			for _ = range ticker.C {
				log.
					Debug().
					Msg("executing jobs")
				jobs.Execute()
			}
		} else {
			jobs.Execute()
		}
	},
}

func init() {
	rootCmd.AddCommand(runCmd)

	runCmd.Flags().IntVarP(&tickSeconds, "tick", "t", 300, "seconds between config checks")
	runCmd.Flags().BoolVarP(&daemon, "daemon", "d", false, "run config checks as a background daemon")
}
