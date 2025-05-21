package main

import (
	"flag"
	"fmt"
	"log"

	"tooling/internal"
)

var (
	projectPath = flag.String("path", "../ManyTests.xcodeproj", "test products path")
	testPlan    = flag.String("testPlan", "", "test plan")
)

func main() {
	flag.Parse()

	if projectPath == nil {
		log.Fatal(fmt.Errorf("incorrect parameters"))
		return
	}

	if err := run(*projectPath, *testPlan); err != nil {
		log.Fatal(err)
	}
}

func run(projectPath, testPlan string) error {
	testProductsPath, err := internal.BuildTestProducts(projectPath, testPlan)
	if err != nil {
		return err
	}

	fmt.Print(testProductsPath)

	return nil
}
