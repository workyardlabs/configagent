package job

import (
	"encoding/json"
	"github.com/rs/zerolog/log"
	"io/ioutil"
)

type Jobs struct {
	Jobs []Job `json:"jobs"`
}

func LoadJobs(filename string) (Jobs, error) {
	jobs := Jobs{}

	jsonStr, err := ioutil.ReadFile(filename)
	if err == nil {
		err = json.Unmarshal(jsonStr, &jobs)
	}

	log.Printf("jobs loaded from file: %s", filename)

	return jobs, err
}

func (j Jobs) Validate() error {
	log.Fatal().Msg("Validate command is not implemented")
	return nil
}

func (j Jobs) Execute() error {
	log.Printf("executing jobs")

	for _, job := range j.Jobs {
		err := job.UpdateConfiguration()
		if err != nil {
			return err
		}
	}

	return nil
}
