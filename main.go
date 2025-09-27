package main

import (
	"flag"
	"fmt"
	"go_transform/transform"
)

func main() {
	help := flag.Bool("help", false, "Display help")
	flag.Parse()

	if *help {
		printHelp()
		return
	}

	x := flag.Args()
	fmt.Print(x)

	err := transform.ExecuteTransform("exampleFiles/transform.json")
	if err != nil {
		panic(err)
	}
}

func printHelp() {
	fmt.Println("Usage: [flags] <transform file path>")
	fmt.Println("Available flags:")
	flag.PrintDefaults()
}
