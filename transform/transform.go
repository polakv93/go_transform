package transform

import (
	"encoding/json"
	"fmt"
	"os"
)

const targetsKey = "__targets"

func ExecuteTransform(transformFilePath string) error {
	transform, _ := readJsonFile(transformFilePath)
	transform, targets := extractTargets(transform)

	for _, target := range targets {
		fileContentToTransform, _ := readJsonFile(target)
		transformed := mergeMaps(fileContentToTransform, transform)
		fmt.Printf("Transformed %s using %s\n", target, transformFilePath)
		out, err := json.MarshalIndent(transformed, "", "  ")
		if err != nil {
			return err
		}
		err = os.WriteFile(target, out, 0644)
		if err != nil {
			return err
		}
	}
	return nil
}

func mergeMaps(base, patch map[string]interface{}) map[string]interface{} {
	for k, v := range patch {
		if bv, ok := base[k].(map[string]interface{}); ok {
			if pv, ok := v.(map[string]interface{}); ok {
				base[k] = mergeMaps(bv, pv)
				continue
			}
		}
		base[k] = v
	}
	return base
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
