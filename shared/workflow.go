package shared

type Workflow struct {
	WorkflowId  string            `yaml:"workflowId"`
	Summary     string            `yaml:"summary,omitempty"`
	Description string            `yaml:"description,omitempty"`
	Inputs      Ref               `yaml:"inputs"`
	Steps       []Step            `yaml:"steps"`
	Outputs     map[string]string `yaml:"outputs,omitempty"`
}
