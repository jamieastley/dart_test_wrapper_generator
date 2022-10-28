package main

import (
	"fmt"
	"os"
	"path/filepath"
)

type FileReader interface {
	getTestFiles(path string) ([]string, error)
	writeFile(output string, path string) error
}

type fileManager struct{}

func (t *fileManager) getTestFiles(path string) ([]string, error) {
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

func (t *fileManager) writeFile(output string, filePath string) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	_, writeErr := file.WriteString(output)
	if writeErr != nil {
		return writeErr
	}

	return nil
}
