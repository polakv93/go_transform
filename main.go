package main

import (
	"flag"
	"fmt"

	"github.com/polakv93/go_transform/transform"
	"github.com/polakv93/go_transform/tui"
)

func main() {
	help := flag.Bool("help", false, "Display help")
	tuiFlag := flag.Bool("tui", false, "Enable the graphical user interface")
	flag.Parse()

	if *tuiFlag {
		if flag.NArg() != 1 {
			printHelp()
			return
		}

		directoryWithTransforms := flag.Arg(0)
		err := tui.RunTui(directoryWithTransforms)
		if err != nil {
			panic(err)
		}
		return
	}

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
	fmt.Println("Usage: <transform file path> or -tui <directory path with transform files> for a graphical user interface")
	fmt.Println("Available flags:")
	flag.PrintDefaults()
}
