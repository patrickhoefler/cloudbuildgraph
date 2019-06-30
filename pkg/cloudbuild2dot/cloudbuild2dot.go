package cloudbuild2dot

import (
	"log"

	"github.com/awalterschulze/gographviz"
	"gopkg.in/yaml.v2"
)

// CloudBuildConfig contains the parts of the Cloud Build configuration schema
// that are relevant for generating the steps graph
type CloudBuildConfig struct {
	Steps []struct {
		ID      string
		WaitFor []string `yaml:"waitFor"`
	}
}

// BuildDotFile builds the Graphviz .dot file from a Google Cloud Build configuration
func BuildDotFile(data []byte) string {
	cloudBuildConfig := CloudBuildConfig{}

	err := yaml.Unmarshal(data, &cloudBuildConfig)
	if err != nil {
		log.Fatal(err)
	}

	if len(cloudBuildConfig.Steps) == 0 {
		return ""
	}

	graph := gographviz.NewEscape()
	graph.SetName("G")
	graph.SetDir(true)
	graph.AddAttr("G", "splines", "ortho")
	graph.AddAttr("G", "nodesep", "0.4")

	startingSteps := []string{
		cloudBuildConfig.Steps[0].ID,
	}

	for index, step := range cloudBuildConfig.Steps {
		graph.AddNode(
			"G",
			step.ID,
			map[string]string{
				"shape": "Mrecord",
				"width": "2",
			},
		)

		for _, waitForStepID := range step.WaitFor {
			if waitForStepID == "-" {
				startingSteps = append(startingSteps, step.ID)
			} else {
				graph.AddEdge(waitForStepID, step.ID, true, nil)
			}
		}

		if step.WaitFor == nil && index > 0 {
			graph.AddEdge(cloudBuildConfig.Steps[index-1].ID, step.ID, true, nil)
		}
	}

	graph.AddSubGraph("G", "rank0", map[string]string{"rank": "same"})

	for _, id := range startingSteps {
		graph.AddNode("rank0", id, nil)
	}

	return graph.String()
}
