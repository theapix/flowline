package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/theapix/flowline/shared"
)

func operationToPlantUML(step shared.Step, operationsMap map[string]shared.Operation) string {
	plantUML := ""
	operation, ok := operationsMap[step.OperationId]

	if ok {
		plantUML += fmt.Sprintf(":%s;\n", operation.Summary)
		plantUML += fmt.Sprintf("note right\n%s\nend note\n", operation.Description)
	}

	return plantUML
}

func workflowToPlantUML(workflow shared.Workflow, operationsMap map[string]shared.Operation, workflowsMap map[string]shared.Workflow, isRoot bool) string {
	var plantUML string
	if !isRoot {
		plantUML = fmt.Sprintf("partition \"%s\" {\n", workflow.Summary)
	}

	for _, step := range workflow.Steps {
		if step.OperationId != "" {
			plantUML += operationToPlantUML(step, operationsMap)
		}
		if step.WorkflowId != "" {
			workflow = workflowsMap[step.WorkflowId]
			plantUML += workflowToPlantUML(workflow, operationsMap, workflowsMap, false)
		}
	}

	if !isRoot {
		plantUML += "}\n"
	}

	return plantUML
}

func flow(config shared.Arazzo, workflowsMap map[string]shared.Workflow, operationsMap map[string]shared.Operation, workflowId string) string {
	workflow := workflowsMap[workflowId]
	plantUML := "@startuml\n"
	plantUML += fmt.Sprintf("title %s:%s [version=%s]\n", config.Info.Title, workflow.Summary, config.Info.Version)
	plantUML += "start\n"

	plantUML += workflowToPlantUML(workflow, operationsMap, workflowsMap, true)

	plantUML += "stop\n"
	plantUML += "@enduml"

	return plantUML
}

func savePlantUML(plantUML string, filePath string) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(plantUML)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	arazzoFilePath := flag.String("arazzo", "", "Arazzo file path")
	workflowId := flag.String("workflow", "", "Workflow id")
	flag.Parse()

	if *arazzoFilePath == "" {
		fmt.Println("Arazzo-file is required")
		os.Exit(1)
	}

	dir := filepath.Dir(*arazzoFilePath)
	config := ArazzoDocument(*arazzoFilePath)
	oasFile := filepath.Join(dir, config.SourceDescriptions[0].Url)
	operationsMap := OperationsMap(oasFile)
	workflowsMap := GetWorkflows(config)

	if *workflowId == "" {
		for id := range workflowsMap {
			savePlantUML(flow(config, workflowsMap, operationsMap, id), fmt.Sprintf("%s-%s[version=%s].puml", config.Info.Title, id, config.Info.Version))
		}
	} else {
		plantUML := flow(config, workflowsMap, operationsMap, *workflowId)
		savePlantUML(plantUML, fmt.Sprintf("%s-%s[version=%s].puml", config.Info.Title, *workflowId, config.Info.Version))
	}

}
