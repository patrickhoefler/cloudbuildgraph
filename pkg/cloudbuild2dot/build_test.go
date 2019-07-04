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

	// valid config
	var validYAMLInput, _ = ioutil.ReadFile("test/valid-input.yaml")
	var validYAMLOutput, _ = ioutil.ReadFile("test/valid-output.dot")

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"valid config",
			args{yamlToCloudBuild(validYAMLInput)},
			string(validYAMLOutput),
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
