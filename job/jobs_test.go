package job

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadJobs(t *testing.T) {
	jobs, err := ImportJobFile(strings.NewReader("invalid"))
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "invalid character")
	assert.Equal(t, Jobs{}, jobs)

	jobs, err = ImportJobFile(strings.NewReader("{,}"))
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "invalid character")
	assert.Equal(t, Jobs{}, jobs)

	jobs, err = ImportJobFile(strings.NewReader("{}"))
	assert.Nil(t, err)
	assert.Equal(t, Jobs{}, jobs)
}
