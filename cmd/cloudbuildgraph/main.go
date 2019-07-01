package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/patrickhoefler/cloudbuildgraph/pkg/cloudbuild2dot"
)

func main() {
	// For now we assume that there is a cloudbuild.yaml file in the working directory
	cloudBuildConfigYAML, err := ioutil.ReadFile("cloudbuild.yaml")
	if err != nil {
		log.SetFlags(0)
		log.Fatal(err)
	}

	dotFileContent := cloudbuild2dot.BuildDotFile(cloudBuildConfigYAML)

	fmt.Println(dotFileContent)
}
