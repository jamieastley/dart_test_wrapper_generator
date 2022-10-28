package main

import "flag"

func parseArgs() *CliArgs {
	path := flag.String("path", "", "Path of the test directory")
	outputFile := flag.String("out", "", "Filename of the output test wrapper")
	template := flag.String("template", "", "(Optional) custom mustache template to use")
	verbose := flag.Bool("v", false, "Enables verbose logging")
	flag.Parse()

	return &CliArgs{
		Path:           *path,
		OutputFilename: *outputFile,
		Template:       *template,
		Verbose:        *verbose,
	}
}

type CliArgs struct {
	Path           string
	OutputFilename string
	Template       string
	Verbose        bool
}
