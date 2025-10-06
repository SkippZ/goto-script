package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	home := os.Getenv("HOME")

	project := os.Args[1]

	project = strings.ToLower(project)
	project = strings.TrimSpace(project)

	replacer := strings.NewReplacer("_", "", "-", "", ".", "")

	project = replacer.Replace(project)

	var path string

	if strings.Contains("kpiservice", project) {
		path = filepath.Join(home, "projects/api/kpi-service/")
	} else if strings.Contains("dpsapi", project) {
		path = filepath.Join(home, "projects/api/dps-api/")
	} else if strings.Contains("exportservice", project) {
		path = filepath.Join(home, "projects/api/export-service/")
	} else if strings.Contains("shopservice", project) {
		path = filepath.Join(home, "projects/api/shop-service/")
	} else if strings.Contains("filestorageservice", project) {
		path = filepath.Join(home, "projects/api/filestorag/")
	} else {
		os.Exit(1)
	}
 
	fmt.Printf("cd %q\n", path)
}
