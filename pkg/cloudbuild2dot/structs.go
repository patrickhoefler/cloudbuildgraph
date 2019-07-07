package cloudbuild2dot

// CloudBuildConfig contains the parts of the Cloud Build configuration schema
// that are relevant for generating the steps graph
type CloudBuildConfig struct {
	Steps []Step
}

// Step contains the parts of a single build step within a Cloud Build configuration
// that are relevant for generating the steps graph
type Step struct {
	ID      string
	Name    string
	Args    []string
	WaitFor []string `yaml:"waitFor"`
}
