package transform

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/jeremywohl/flatten"
)

import "github.com/tidwall/sjson"

const targetsKey = "__targets"

func ExecuteTransform(transformFilePath string) error {
	transform, err := readJsonFile(transformFilePath)
	if err != nil {
		return err
	}
	transform, targets := extractTargets(transform)
	flattenTransformKeys, err := flatten.Flatten(transform, "", flatten.DotStyle)
	if err != nil {
		panic(fmt.Sprintf("failed to flatten transform: %v", err))
	}

	for _, target := range targets {
		fileContentBytes, _ := os.ReadFile(target)

		for path := range flattenTransformKeys {
			fileContentBytes, _ = sjson.SetBytes(fileContentBytes, path, flattenTransformKeys[path])
		}

		fmt.Printf("Transformed %s using %s\n", target, transformFilePath)
		err = os.WriteFile(target, fileContentBytes, 0644)
		if err != nil {
			return err
		}
	}
	return nil
}

func extractTargets(transform map[string]interface{}) (map[string]interface{}, []string) {
	targets, exists := transform[targetsKey]
	if !exists {
		panic(fmt.Sprintf("%s not found in transform", targetsKey))
	}
	delete(transform, targetsKey)

	listOfTargets := targets.([]interface{})
	ret := make([]string, len(listOfTargets))
	for i, target := range listOfTargets {
		ret[i] = target.(string)
	}

	return transform, ret
}

func readJsonFile(fileName string) (map[string]interface{}, error) {
	transformData, err := os.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	var transform map[string]interface{}
	err = json.Unmarshal(transformData, &transform)
	if err != nil {
		return nil, err
	}
	return transform, nil
}
