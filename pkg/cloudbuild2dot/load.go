package cloudbuild2dot

import (
	"errors"
	"log"
	"os"

	"github.com/spf13/afero"
)

// CloudBuildConfig contains the parts of the Cloud Build configuration schema
// that are relevant for generating the steps graph
type CloudBuildConfig struct {
	Steps []CloudBuildConfigStep
}

// CloudBuildConfigStep represents a single build step within a Cloud Build configuration
type CloudBuildConfigStep struct {
	ID      string
	WaitFor []string `yaml:"waitFor"`
}

// LoadCloudBuildConfig looks for the Cloud Build configuration file (currently
// only YAML is supported) and returns a CloudBuildConfig.
func LoadCloudBuildConfig() (cloudBuildConfig CloudBuildConfig, err error) {
	return loadYAMLConfig(afero.NewOsFs())
}

func loadYAMLConfig(AppFs afero.Fs) (CloudBuildConfig, error) {
	for _, filename := range []string{"cloudbuild.yaml", "cloudbuild.yml"} {
		yamlContent, err := afero.ReadFile(AppFs, filename)
		if err != nil {
			if os.IsNotExist(err) {
				continue
			} else {
				log.Fatal(err)
			}
		}

		return yamlToCloudBuild(yamlContent), nil
	}

	return CloudBuildConfig{}, errors.New("Could not find cloudbuild.y(a)ml in the current working directory")
}
