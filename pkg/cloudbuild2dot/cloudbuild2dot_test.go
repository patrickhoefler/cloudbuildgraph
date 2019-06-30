package cloudbuild2dot_test

import (
	"fmt"
	"io/ioutil"

	"github.com/patrickhoefler/cloudbuildgraph/pkg/cloudbuild2dot"
)

func ExampleBuildDotFile() {
	data, _ := ioutil.ReadFile("test/valid-cloud-build-config.yaml")
	fmt.Println(cloudbuild2dot.BuildDotFile(data))
	// Output:
	// digraph G {
	// 	nodesep=0.4;
	// 	splines=ortho;
	// 	"step-a"->"step-b";
	// 	"step-a"->"step-c";
	// 	subgraph rank0 {
	// 	rank=same;
	// 	"step-a" [ shape=Mrecord, width=2 ];
	// 	"step-d" [ shape=Mrecord, width=2 ];
	//
	// }
	// ;
	// 	"step-b" [ shape=Mrecord, width=2 ];
	// 	"step-c" [ shape=Mrecord, width=2 ];
	//
	// }
}
