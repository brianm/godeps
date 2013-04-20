package main

import (
	"go/build"
	"log"
	"strings"
)

func main() {
	ctx := build.Default
	visited := make(map[string]bool)

	pkg, err := ctx.Import("gitlab.ningeng.net/brianm/gohouse", ".", build.AllowBinary)
	if err != nil {
		log.Fatalf("error on import: %s", err);
	}			
	explort(ctx, pkg, visited)	
}

func explort(ctx build.Context, pkg *build.Package, visited map[string]bool) {
	for _, packageName := range pkg.Imports {
		if !visited[packageName] {
			visited[packageName] = true
			if strings.Contains(packageName, ".") {
				log.Printf("%s\n", packageName)
			}
			if ! (packageName == "C") {
				child, err := ctx.Import(packageName, ".", build.AllowBinary)
				if err != nil {
					log.Fatalf("error on import: %s", err);
				}					
				explort(ctx, child, visited)
			}
		}
	}
}