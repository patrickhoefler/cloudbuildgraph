package cloudbuild2dot

import (
	"io/ioutil"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestBuildDotFile(t *testing.T) {
	type args struct {
		cloudBuildConfig CloudBuildConfig
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"valid config",
			args{yamlToCloudBuild(loadTestFile("valid-input.yaml"))},
			string(loadTestFile("valid-output.dot")),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := BuildDotFile(tt.args.cloudBuildConfig)

			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("BuildDotFile() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func loadTestFile(filename string) []byte {
	testFileContent, _ := ioutil.ReadFile("test/" + filename)

	return testFileContent
}
