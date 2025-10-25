// cmd/gcommit/main.go
package main

import (
	"fmt"
	"gcommit/internal/generator"
	"gcommit/internal/git"
	"log"
)

func main() {
	files, err := git.GetChangedFiles()
	if err != nil {
		log.Fatal(err)
	}

	scope, scopeList := generator.DetectScopeWithList(files)
	fmt.Printf("Detected scope: %s\n", scope, scopeList)
}
