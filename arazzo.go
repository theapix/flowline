package main

import (
	"log"
	"os"

	"github.com/theapix/flowline/shared"
	"gopkg.in/yaml.v3"
)

func GetWorkflows(config shared.Arazzo) map[string]shared.Workflow {
	workflowsMap := make(map[string]shared.Workflow)
	for _, workflow := range config.Workflows {
		workflowsMap[workflow.WorkflowId] = workflow
	}

	return workflowsMap
}

func ArazzoDocument(filePath string) shared.Arazzo {
	data, err := os.ReadFile(filePath)

	if err != nil {
		log.Fatalf("error: %v", err)
	}

	var config shared.Arazzo
	err = yaml.Unmarshal(data, &config)

	if err != nil {
		log.Fatalf("error: %v", err)
	}

	return config
}
