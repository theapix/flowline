package main

import (
	"log"
	"os"

	"github.com/theapix/flowline/shared"
	"gopkg.in/yaml.v3"
)

type OpenAPISpec struct {
	Paths map[string]map[string]shared.Operation `yaml:"paths"`
}

func findOperations(filePath string) (map[string]shared.Operation, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var spec OpenAPISpec
	err = yaml.Unmarshal(data, &spec)
	if err != nil {
		return nil, err
	}

	operationsMap := make(map[string]shared.Operation)

	for _, methods := range spec.Paths {
		for _, operation := range methods {
			if operation.OperationID != "" {
				operationsMap[operation.OperationID] = operation
			}
		}
	}

	return operationsMap, nil
}

func OperationsMap(filePath string) map[string]shared.Operation {
	operationsMap, err := findOperations(filePath)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	return operationsMap
}
