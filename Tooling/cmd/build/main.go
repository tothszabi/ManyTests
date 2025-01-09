package main

import (
	"flag"
	"fmt"
	"log"

	"tooling/internal"
)

var (
	projectPath = flag.String("path", "../ManyTests.xcodeproj", "test products path")
)

func main() {
	flag.Parse()

	if projectPath == nil {
		log.Fatal(fmt.Errorf("incorrect parameters"))
		return
	}

	if err := run(*projectPath); err != nil {
		log.Fatal(err)
	}
}

func run(projectPath string) error {
	testProductsPath, err := internal.BuildTestProducts(projectPath)
	if err != nil {
		return err
	}

	fmt.Print(testProductsPath)

	return nil
}
