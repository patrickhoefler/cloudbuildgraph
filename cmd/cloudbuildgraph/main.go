package main

import (
	"fmt"
	"log"

	"github.com/patrickhoefler/cloudbuildgraph/pkg/cloudbuild2dot"
)

func main() {
	log.SetFlags(0)

	cloudBuildConfig, err := cloudbuild2dot.LoadCloudBuildConfig()
	if err != nil {
		log.Fatal(err)
	}

	dotFileContent := cloudbuild2dot.BuildDotFile(cloudBuildConfig)

	fmt.Print(dotFileContent)
}
