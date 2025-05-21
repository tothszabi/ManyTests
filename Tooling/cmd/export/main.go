package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"

	"tooling/internal"
)

var (
	testProductsPath = flag.String("path", "../test-products.xctestproducts", "test products path")
	testPlan         = flag.String("testPlan", "", "test plan")
	numBuckets       = flag.Int("buckets", 5, "number of buckets")
	key              = flag.String("key", "TEST_SHARD", "test shard key")
)

func main() {
	flag.Parse()

	if testProductsPath == nil || numBuckets == nil || key == nil {
		log.Fatal(fmt.Errorf("incorrect parameters"))
		return
	}

	if err := run(*testProductsPath, *testPlan, *numBuckets); err != nil {
		log.Fatal(err)
	}
}

func run(testProductsPath, testPlan string, numBuckets int) error {
	tests, err := internal.CollectTests(testProductsPath, testPlan)
	if err != nil {
		return err
	}

	buckets := bucketTests(tests, numBuckets)

	if err := export(buckets, *key); err != nil {
		return err
	}

	return nil
}

func bucketTests(tests []string, numBuckets int) [][]string {
	buckets := make([][]string, numBuckets)
	bucketSize := (len(tests) + numBuckets - 1) / numBuckets

	for i, test := range tests {
		bucketIndex := i / bucketSize
		buckets[bucketIndex] = append(buckets[bucketIndex], test)
	}

	return buckets
}

func export(buckets [][]string, key string) error {
	tmpDir, err := internal.CreateTempFolder()
	if err != nil {
		return err
	}

	tmpDir = "/Users/szabi/Developer/misc/ManyTests/test-output"

	for i, bucket := range buckets {
		bytes, err := json.Marshal(bucket)
		if err != nil {
			return err
		}

		filePath := filepath.Join(tmpDir, fmt.Sprintf("test_shard_%d.txt", i))
		err = os.WriteFile(filePath, bytes, 0644)
		if err != nil {
			return err
		}

		fmt.Println(filePath)

		numberedKey := fmt.Sprintf("%s_%d", key, i)
		cmd := exec.Command("envman", "add", "--key", numberedKey, "--value", filePath)
		if err := internal.Execute(cmd); err != nil {
			return err
		}
	}

	return nil
}
