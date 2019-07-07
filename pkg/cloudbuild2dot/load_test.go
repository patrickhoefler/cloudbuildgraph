package cloudbuild2dot

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/spf13/afero"
)

func Test_loadConfig(t *testing.T) {
	type args struct {
		AppFs afero.Fs
	}

	yamlFileContent := []byte("{steps: [{id: yaml}]}")
	ymlFileContent := []byte("{steps: [{id: yml}]}")
	jsonFileContent := []byte("{\"steps\": [{\"id\": \"json\"}]}")

	// cloudbuild.yaml
	yamlFS := afero.NewMemMapFs()
	afero.WriteFile(yamlFS, "cloudbuild.yaml", yamlFileContent, 0644)
	afero.WriteFile(yamlFS, "cloudbuild.yml", ymlFileContent, 0644)
	afero.WriteFile(yamlFS, "cloudbuild.json", jsonFileContent, 0644)

	// cloudbuild.yml
	ymlFS := afero.NewMemMapFs()
	afero.WriteFile(ymlFS, "cloudbuild.yml", ymlFileContent, 0644)
	afero.WriteFile(ymlFS, "cloudbuild.json", jsonFileContent, 0644)

	// cloudbuild.json
	jsonFS := afero.NewMemMapFs()
	afero.WriteFile(jsonFS, "cloudbuild.json", jsonFileContent, 0644)

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
		{
			"cloudbuild.json",
			args{jsonFS},
			jsonToCloudBuild(jsonFileContent),
			false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := loadConfig(tt.args.AppFs)

			if (err != nil) != tt.wantErr {
				t.Errorf("loadConfig() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("loadConfig() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
