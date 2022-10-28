package main

import (
	"fmt"
	"os"
	"path/filepath"
)

type FileReader interface {
	getTestFiles(path string) ([]string, error)
}

type fileReader struct{}

func (t *fileReader) getTestFiles(path string) ([]string, error) {
	var files []string
	fileErr := filepath.Walk(path, func(currentPath string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		fmt.Println(currentPath)
		files = append(files, currentPath)
		return nil
	})

	return files, fileErr
}
