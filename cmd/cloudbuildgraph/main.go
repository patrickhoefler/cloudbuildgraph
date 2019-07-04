package main

import (
	"fmt"

	"github.com/patrickhoefler/cloudbuildgraph/pkg/cloudbuild2dot"
)

func main() {
	cloudBuildConfig, _ := cloudbuild2dot.LoadCloudBuildConfig()

	dotFileContent := cloudbuild2dot.BuildDotFile(cloudBuildConfig)

	fmt.Print(dotFileContent)
}
