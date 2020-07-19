package job

import (
	"github.com/rs/zerolog/log"
	"github.com/workyard/configagent/parameter"
)

type ParameterKey struct {
	Key  string `json:"key"`
	Path string `json:"path"`
}

type Job struct {
	Name          string         `json:"name"`
	ParameterKeys []ParameterKey `json:"parameterKeys"`
	ConfigFiles   []ConfigFile   `json:"configFiles"`
	Restart       bool           `json:"restartOnChange"`
	Service       string         `json:"service"`
}

func (j Job) UpdateConfiguration() error {
	params, err := PopulateParameterMap(j.ParameterKeys, parameter.SSMParameterStore{})
	if err != nil {
		return err
	}

	// A service needs to be restarted if any of it's configuration
	// files has been modified.
	var restartRequired bool = false

	for _, config := range j.ConfigFiles {
		fileChanged, err := config.UpdateExistingFile(params)
		if err != nil {
			return err
		}

		if fileChanged {
			restartRequired = true
		}
	}

	if restartRequired && j.Restart {
		err = j.RestartService()
		if err != nil {
			return err
		}
	}
	return nil
}

func (j Job) RestartService() error {
	log.Info().Msgf("TODO:  restart serice %s", j.Service)
	return nil
}
