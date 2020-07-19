package job

import (
	"encoding/json"
	"github.com/rs/zerolog/log"
	"io"
	"io/ioutil"
)

type Jobs struct {
	Jobs []Job `json:"jobs"`
}

func ImportJobFile(rdr io.Reader) (Jobs, error) {
	jobs := Jobs{}

	jsonStr, err := ioutil.ReadAll(rdr)
	if err == nil {
		err = json.Unmarshal(jsonStr, &jobs)
	}

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
