package shared

type Step struct {
	StepId      string            `yaml:"stepId"`
	OperationId string            `yaml:"operationId,omitempty"`
	WorkflowId  string            `yaml:"workflowId,omitempty"`
	Description string            `yaml:"description,omitempty"`
	Outputs     map[string]string `yaml:"outputs,omitempty"`
}
