package cloudbuild2dot

import (
	"errors"
	"log"
	"os"

	"github.com/spf13/afero"
)

// LoadCloudBuildConfig looks for the Cloud Build configuration file (currently
// only YAML is supported) and returns a CloudBuildConfig.
func LoadCloudBuildConfig() (cloudBuildConfig CloudBuildConfig, err error) {
	return loadConfig(afero.NewOsFs())
}

func loadConfig(AppFs afero.Fs) (CloudBuildConfig, error) {
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

	for _, filename := range []string{"cloudbuild.json"} {
		jsonContent, err := afero.ReadFile(AppFs, filename)
		if err != nil {
			if os.IsNotExist(err) {
				continue
			} else {
				log.Fatal(err)
			}
		}

		return jsonToCloudBuild(jsonContent), nil
	}

	return CloudBuildConfig{}, errors.New("Could not find any Cloud Build config file in the current working directory")
}
