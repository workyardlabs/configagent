package job

import (
	"github.com/rs/zerolog/log"
	"github.com/workyard/configagent/parameter"
)

type ParamMap map[string]string

func PopulateParameterMap(keys []ParameterKey, store parameter.IParameterStore) (ParamMap, error) {
	params := make(ParamMap)

	for _, key := range keys {
		value, err := store.Load(key.Path)
		if err != nil {
			return make(map[string]string), err
		}

		params[key.Key] = value

		log.Printf("loading config %s", key.Key)
	}

	return params, nil
}
