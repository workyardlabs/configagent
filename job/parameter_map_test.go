package job

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

type MockParameterStore struct{}

func (m MockParameterStore) Load(path string) (string, error) {
	if path == "name" {
		return "Tristan Wontonobee", nil
	} else if path == "backup" {
		return "true", nil
	} else {
		return "", errors.New("path does not exist")
	}
}

func TestPopulateParameterMap(t *testing.T) {
	store := MockParameterStore{}

	// empty params keys
	keys := []ParameterKey{}
	params, err := PopulateParameterMap(keys, store)
	assert.Empty(t, params)
	assert.Nil(t, err)

	// list of valid keys
	keys = []ParameterKey{
		ParameterKey{Key: "name", Path: "name"},
		ParameterKey{Key: "backup", Path: "backup"},
	}

	params, err = PopulateParameterMap(keys, store)
	assert.Equal(t, "Tristan Wontonobee", params["name"])
	assert.Equal(t, "true", params["backup"])
	assert.Equal(t, 2, len(params))
	assert.Nil(t, err)

	// has an invalid key
	keys = []ParameterKey{
		ParameterKey{Key: "name", Path: "name"},
		ParameterKey{Key: "backup", Path: "doesnotexist"},
		ParameterKey{Key: "backup", Path: "backup"},
	}

	params, err = PopulateParameterMap(keys, store)
	assert.Empty(t, params)
	assert.NotNil(t, err)
}
