package main

import "flag"

func parseArgs() *CliArgs {
	path := flag.String("path", "", "Path of the test directory")
	outputFile := flag.String("out", "", "Filename of the output test wrapper")
	verbose := flag.Bool("v", false, "Enables verbose logging")
	flag.Parse()

	return &CliArgs{
		Path:       *path,
		OutputFile: *outputFile,
		Verbose:    *verbose,
	}
}

type CliArgs struct {
	Path       string
	OutputFile string
	Verbose    bool
}
