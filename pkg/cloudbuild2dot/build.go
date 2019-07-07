package cloudbuild2dot

import (
	"strings"

	"github.com/awalterschulze/gographviz"
)

// BuildDotFile builds a GraphViz .dot file from a Google Cloud Build configuration
func BuildDotFile(cloudBuildConfig CloudBuildConfig) string {
	graph := gographviz.NewEscape()
	graph.SetName("G")
	graph.SetDir(true)
	graph.AddAttr("G", "splines", "ortho")
	graph.AddAttr("G", "nodesep", "0.4")

	startingSteps := []string{
		getStepLabel(cloudBuildConfig.Steps[0]),
	}

	for index, step := range cloudBuildConfig.Steps {
		graph.AddNode(
			"G",
			getStepLabel(step),
			map[string]string{
				"shape": "Mrecord",
				"width": "2",
			},
		)

		for _, waitForStepID := range step.WaitFor {
			if waitForStepID == "-" {
				startingSteps = append(startingSteps, getStepLabel(step))
			} else {
				graph.AddEdge(
					waitForStepID,
					getStepLabel(step),
					true,
					nil,
				)
			}
		}

		if step.WaitFor == nil && index > 0 {
			graph.AddEdge(
				getStepLabel(cloudBuildConfig.Steps[index-1]),
				getStepLabel(step),
				true,
				nil,
			)
		}
	}

	graph.AddSubGraph("G", "rank0", map[string]string{"rank": "same"})

	for _, id := range startingSteps {
		graph.AddNode("rank0", id, nil)
	}

	return graph.String()
}

func getStepLabel(step Step) string {
	if step.ID != "" {
		return step.ID
	}

	splitName := strings.Split(step.Name, "/")

	if len(step.Args) > 0 {
		return splitName[len(splitName)-1] + " " + step.Args[0]
	}

	return splitName[len(splitName)-1]
}
