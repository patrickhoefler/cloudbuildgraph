package cloudbuild2dot

import (
	"log"

	"gopkg.in/yaml.v2"
)

func yamlToCloudBuild(yamlContent []byte) CloudBuildConfig {
	cloudBuildConfig := CloudBuildConfig{}

	err := yaml.Unmarshal(yamlContent, &cloudBuildConfig)
	if err != nil {
		log.Fatal(err)
	}

	return cloudBuildConfig
}
