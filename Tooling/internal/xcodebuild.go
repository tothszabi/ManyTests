package internal

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

const (
	testProductsName = "test-products.xctestproducts"
)

func BuildTestProducts(projectPath string) (string, error) {
	tmpDir, err := CreateTempFolder()
	if err != nil {
		return "", err
	}

	testProductsPath := filepath.Join(tmpDir, testProductsName)
	cmd := exec.Command("xcodebuild", "build-for-testing", "-testProductsPath", testProductsPath, "-project", projectPath, "-scheme", "ManyTests", "-destination", "platform=iOS Simulator,name=iPhone 16 Pro Max,OS=latest")

	if err := Execute(cmd); err != nil {
		return "", err
	}

	return testProductsPath, nil
}

func CollectTests(testProductsPath string) ([]string, error) {
	tmpDir, err := CreateTempFolder()
	if err != nil {
		return nil, err
	}

	testOutput := filepath.Join(tmpDir, "result.txt")
	cmd := exec.Command("xcodebuild", "test-without-building", "-enumerate-tests", "-test-enumeration-format", "json", "-test-enumeration-style", "flat", "-test-enumeration-output-path", testOutput, "-testProductsPath", testProductsPath, "-destination", "platform=iOS Simulator,name=iPhone 16 Pro Max,OS=latest")

	if err := Execute(cmd); err != nil {
		return nil, err
	}

	bytes, err := os.ReadFile(testOutput)
	if err != nil {
		return nil, err
	}

	type testData struct {
		Values []struct {
			Tests []struct {
				Identifier string `json:"identifier"`
			} `json:"enabledTests"`
		} `json:"values"`
	}

	var data testData
	if err := json.Unmarshal(bytes, &data); err != nil {
		return nil, err
	}

	var tests []string
	for _, value := range data.Values {
		for _, test := range value.Tests {
			tests = append(tests, test.Identifier)
		}
	}

	if err := os.Remove(testOutput); err != nil {
		return nil, err
	}

	return tests, nil
}

func RunTests(testProductsPath, testShard string) error {
	bytes, err := os.ReadFile(testShard)
	if err != nil {
		return err
	}

	var testNames []string
	if err := json.Unmarshal(bytes, &testNames); err != nil {
		return err
	}

	var testArg string
	for _, testName := range testNames {
		testArg += fmt.Sprintf("-only-testing:%s ", testName)
	}

	cmd := exec.Command("xcodebuild", "test-without-building", "-testProductsPath", testProductsPath, testArg, "-destination", "platform=iOS Simulator,name=iPhone 16 Pro Max,OS=latest")
	cmd.Stdout = os.Stdout

	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}
