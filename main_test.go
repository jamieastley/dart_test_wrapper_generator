package main

import (
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

type mockFileManager struct {
	TestFiles []string
	Error     error
}

func (m *mockFileManager) getTestFiles(_ string) ([]string, error) {
	return m.TestFiles, m.Error
}

func (m *mockFileManager) writeFile(_ string, _ string) error {
	return m.Error
}

type MockLogger struct{}

func (m MockLogger) Info(args ...interface{}) {
	fmt.Println(args...)
}

func (m MockLogger) Debug(args ...interface{}) {
	fmt.Println(args...)
}

func (m MockLogger) Error(args ...interface{}) {
	fmt.Println(args...)
}

func TestParseTestFiles(t *testing.T) {

	var tests = []struct {
		mockFileManager FileManager
		args            CliArgs
		importResults   []DartImport
		error           error
	}{
		{
			mockFileManager: &mockFileManager{
				TestFiles: []string{
					"models/model_helper.dart",
					"models/model_test.dart",
					"models/another_model_test.dart",
				},
				Error: nil,
			},
			args: CliArgs{
				Path:           "/home",
				OutputFilename: "wrapper",
			},
			importResults: []DartImport{
				{
					Alias: "model_test",
					Path:  "models/model_test.dart",
				},
				{
					Alias: "another_model_test",
					Path:  "models/another_model_test.dart",
				},
			},
		},
		{
			mockFileManager: &mockFileManager{
				TestFiles: []string{},
				Error:     nil,
			},
			args: CliArgs{
				Path:           "/home",
				OutputFilename: "wrapper",
			},
			importResults: []DartImport{},
			error:         errors.New("no files found"),
		},
		{
			mockFileManager: &mockFileManager{
				TestFiles: []string{
					"file.txt",
					"README.md",
				},
			},
			args: CliArgs{
				Path:           "/home",
				OutputFilename: "wrapper",
			},
			importResults: []DartImport{},
			error:         errors.New("no test files found"),
		},
	}

	for _, test := range tests {
		result, err := ParseTestFiles(test.mockFileManager, &test.args, &MockLogger{})
		assert.Equal(t, test.importResults, result)
		assert.Equal(t, test.error, err)
	}
}
