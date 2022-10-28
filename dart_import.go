package main

type DartImport struct {
	Alias string
	Path  string
}

func CreateMustacheData(imports []DartImport) map[string]*[]DartImport {
	return map[string]*[]DartImport{
		"importList": &imports,
	}
}
