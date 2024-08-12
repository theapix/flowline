package shared

type Operation struct {
	OperationID string `yaml:"operationId"`
	Summary     string `yaml:"summary"`
	Description string `yaml:"description"`
}
