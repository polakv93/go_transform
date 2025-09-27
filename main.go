package main

import (
	"flag"
	"fmt"
	"github.com/polakv93/go_transform/transform"
)

func main() {
	help := flag.Bool("help", false, "Display help")
	flag.Parse()

	if *help || flag.NArg() != 1 {
		printHelp()
		return
	}

	transformFilePath := flag.Arg(0)

	err := transform.ExecuteTransform(transformFilePath)
	if err != nil {
		panic(err)
	}
}

func printHelp() {
	fmt.Println("Usage: [flags] <transform file path>")
	fmt.Println("Available flags:")
	flag.PrintDefaults()
}
