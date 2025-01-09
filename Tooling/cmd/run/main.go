package main

import (
	"flag"
	"fmt"
	"log"

	"tooling/internal"
)

var (
	testProductsPath = flag.String("path", "../test-products.xctestproducts", "test products path")
	testShard        = flag.String("shard", "", "test shard")
)

func main() {
	flag.Parse()

	if testProductsPath == nil || testShard == nil || *testShard == "" {
		log.Fatal(fmt.Errorf("incorrect parameters"))
		return
	}

	if err := run(*testProductsPath, *testShard); err != nil {
		log.Fatal(err)
	}
}

func run(testProductsPath, testShard string) error {
	if err := internal.RunTests(testProductsPath, testShard); err != nil {
		return err
	}

	return nil
}
