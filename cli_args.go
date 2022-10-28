package main

import "flag"

func parseArgs() *CliArgs {
	path := flag.String("path", "", "Path of the test directory")
	flag.Parse()

	return &CliArgs{
		Path: *path,
	}
}

type CliArgs struct {
	Path string
}
