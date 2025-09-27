package main

import (
	"go_transform/transform"
)

func main() {
	err := transform.ExecuteTransform("exampleFiles/transform.json")
	if err != nil {
		panic(err)
	}
}
