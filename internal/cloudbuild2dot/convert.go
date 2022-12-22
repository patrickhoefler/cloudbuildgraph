package cloudbuild2dot

import (
	"encoding/json"
	"log"

	"gopkg.in/yaml.v3"
)

func jsonToCloudBuild(content []byte) CloudBuildConfig {
	cloudBuildConfig := CloudBuildConfig{}

	err := json.Unmarshal(content, &cloudBuildConfig)
	if err != nil {
		log.Fatal(err)
	}

	return cloudBuildConfig
}

func yamlToCloudBuild(content []byte) CloudBuildConfig {
	cloudBuildConfig := CloudBuildConfig{}

	err := yaml.Unmarshal(content, &cloudBuildConfig)
	if err != nil {
		log.Fatal(err)
	}

	return cloudBuildConfig
}
