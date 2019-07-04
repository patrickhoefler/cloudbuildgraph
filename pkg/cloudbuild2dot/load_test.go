package cloudbuild2dot

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/spf13/afero"
)

func Test_loadYAMLConfig(t *testing.T) {
	type args struct {
		AppFs afero.Fs
	}

	var yamlFileContent = []byte("{steps: [{id: yaml}]}")
	var ymlFileContent = []byte("{steps: [{id: yml}]}")

	// cloudbuild.yaml
	var yamlFS = afero.NewMemMapFs()
	afero.WriteFile(yamlFS, "cloudbuild.yaml", yamlFileContent, 0644)
	afero.WriteFile(yamlFS, "cloudbuild.yml", ymlFileContent, 0644)

	// cloudbuild.yml
	var ymlFS = afero.NewMemMapFs()
	afero.WriteFile(ymlFS, "cloudbuild.yml", ymlFileContent, 0644)

	tests := []struct {
		name    string
		args    args
		want    CloudBuildConfig
		wantErr bool
	}{
		{
			"No YAML files",
			args{afero.NewMemMapFs()},
			CloudBuildConfig{},
			true,
		},
		{
			"cloudbuild.yaml",
			args{yamlFS},
			yamlToCloudBuild(yamlFileContent),
			false,
		},
		{
			"cloudbuild.yml",
			args{ymlFS},
			yamlToCloudBuild(ymlFileContent),
			false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := loadYAMLConfig(tt.args.AppFs)

			if (err != nil) != tt.wantErr {
				t.Errorf("loadYAMLConfig() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("loadYAMLConfig() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
