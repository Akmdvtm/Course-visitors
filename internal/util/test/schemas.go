package test

import (
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

func GetDeploymentSchemas() []string {
	_, file, _, ok := runtime.Caller(0)
	if !ok {
		log.Fatalf("Failed to get source file path")
	}
	dir := filepath.Join(filepath.Dir(file), "..", "..", "..", "deployment", "sql")
	files, err := os.ReadDir(dir)
	if err != nil {
		log.Fatalf("Failed to list files in %s: %v", dir, err)
	}

	schemas := make([]string, 0, len(files))
	patches := make([]string, 0, len(files))
	for i := range files {
		f := files[i]
		fullPath := filepath.Join(dir, f.Name())

		if f.IsDir() || filepath.Ext(fullPath) != ".sql" {
			continue
		}

		if strings.HasPrefix(f.Name(), "init") {
			schemas = append(schemas, fullPath)
		}
		if strings.HasPrefix(f.Name(), "patch") {
			patches = append(patches, fullPath)
		}
	}

	return append(schemas, patches...)
}
