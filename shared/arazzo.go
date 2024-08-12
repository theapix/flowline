package shared

// Define a struct that matches the structure of your YAML file
type Arazzo struct {
	Info struct {
		Title   string `yaml:"title"`
		Version string `yaml:"version"`
	} `yaml:"info"`
	SourceDescriptions []SourceDescriptions `yaml:"sourceDescriptions,omitempty"`
	Workflows          []Workflow           `yaml:"workflows"`
}
