package cloudbuild2dot_test

import (
	"io/ioutil"
	"testing"

	"github.com/patrickhoefler/cloudbuildgraph/pkg/cloudbuild2dot"
	"github.com/stretchr/testify/assert"
)

func TestValidCloudBuildConfig(t *testing.T) {
	validInput, _ := ioutil.ReadFile("test/valid-input.yaml")
	validOutput, _ := ioutil.ReadFile("test/valid-output.dot")
	assert.Equal(
		t,
		string(validOutput),
		cloudbuild2dot.BuildDotFile(validInput),
	)
}

func TestInputFileNotFound(t *testing.T) {
	cloudBuildConfig := cloudbuild2dot.BuildDotFile([]byte(""))
	assert.Equal(
		t,
		"",
		cloudBuildConfig,
	)
}
