package job

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"text/template"
	"time"

	"github.com/rs/zerolog/log"
)

type ConfigFile struct {
	Backup       bool   `json:"backupConfig"`
	TemplatePath string `json:"templatePath"`
	ConfigPath   string `json:"configPath"`
}

func (cf ConfigFile) UpdateExistingFile(params ParamMap) (bool, error) {
	var fileChanged bool = false

	configBytes, err := cf.GenerateFromTemplate(params)
	if err != nil {
		return false, err
	}

	existing, _ := ioutil.ReadFile(cf.ConfigPath)
	equal := CompareBytes(configBytes, existing)

	if !equal {
		err := cf.Write(configBytes)
		if err != nil {
			return false, err
		}

		fileChanged = true
	}

	return fileChanged, nil
}

func (cf ConfigFile) Write(content []byte) error {
	err := cf.BackupConfig()
	if err != nil {
		return err
	}

	log.Info().Msgf("writing config file %s (%d bytes)", cf.ConfigPath, len(content))
	fileErr := ioutil.WriteFile(cf.ConfigPath, content, 0644)

	return fileErr
}

func (cf ConfigFile) BackupConfig() error {
	if cf.Backup {
		if _, err := os.Stat(cf.ConfigPath); err == nil {
			t := time.Now()

			backupFilename := fmt.Sprintf("%s.%d", cf.ConfigPath, t.Unix())
			log.Printf("writing backup file %s", backupFilename)

			bytes, err := ioutil.ReadFile(cf.ConfigPath)
			if err != nil {
				return err
			}

			err = ioutil.WriteFile(backupFilename, bytes, 0644)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (cf ConfigFile) GenerateFromTemplate(params ParamMap) ([]byte, error) {
	tpl, err := ioutil.ReadFile(cf.TemplatePath)
	if err != nil {
		return []byte{}, err
	}

	t, err := template.New("t").Parse(string(tpl))
	if err != nil {
		return []byte{}, err
	}

	var b bytes.Buffer
	writer := bufio.NewWriter(&b)
	t.Execute(writer, params)
	writer.Flush()

	return b.Bytes(), nil
}
