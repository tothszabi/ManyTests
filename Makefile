all-tests:
	xcodebuild \
	test \
	-scheme ManyTests \
	-destination 'platform=iOS Simulator,name=iPhone 16 Pro Max,OS=latest' \

list-tests:
	xcodebuild \
	test \
	-enumerate-tests \
	-test-enumeration-format json \
	-test-enumeration-style flat \
	-test-enumeration-output-path /Users/szabi/Developer/misc/ManyTests/Tooling/result.txt \
	-scheme ManyTests \
	-destination 'platform=iOS Simulator,name=iPhone 16 Pro Max,OS=latest' \

target-only-tests:
	xcodebuild \
	test \
	-only-testing ManyTestsTests \
	-scheme ManyTests \
	-destination 'platform=iOS Simulator,name=iPhone 16 Pro Max,OS=latest' \

class-only-tests:
	xcodebuild \
	test \
	-only-testing ManyTestsTests/ManyTestsTests \
	-scheme ManyTests \
	-destination 'platform=iOS Simulator,name=iPhone 16 Pro Max,OS=latest' \

function-only-tests:
	xcodebuild \
	test \
	-only-testing ManyTestsTests/ManyTestsTests/testAnotherExample \
	-scheme ManyTests \
	-destination 'platform=iOS Simulator,name=iPhone 16 Pro Max,OS=latest' \

multiple-function-tests:
	xcodebuild \
	test \
	-only-testing:ManyTestsTests/ManyTestsTests/testAnotherExample \
	-only-testing:ManyTestsUITests/ManyTestsUITests/testAnotherExample \
	-only-testing:ManyTestsUITests/ManyTestsUITestsLaunchTests/testAnotherLaunch \
	-scheme ManyTests \
	-destination 'platform=iOS Simulator,name=iPhone 16 Pro Max,OS=latest' \

mixed-tests:
	xcodebuild \
	test \
	-only-testing:ManyTestsTests \
	-only-testing:ManyTestsUITests/ManyTestsUITests/testAnotherExample \
	-scheme ManyTests \
	-destination 'platform=iOS Simulator,name=iPhone 16 Pro Max,OS=latest' \

build-tests:
	xcodebuild \
	build-for-testing \
	-scheme ManyTests \
	-destination 'platform=iOS Simulator,name=iPhone 16 Pro Max,OS=latest' \
	-testProductsPath ./test-products \

list-prebuilt-tests:
	xcodebuild \
	test-without-building \
	-enumerate-tests \
	-test-enumeration-format json \
	-test-enumeration-style flat \
	-testProductsPath test-products.xctestproducts \
	-destination 'platform=iOS Simulator,name=iPhone 16 Pro Max,OS=latest' \

run-subset-prebuilt-tests:
	xcodebuild \
	test-without-building \
	-testProductsPath test-products.xctestproducts \
	-only-testing:ManyTestsTests/ManyTestsTests/testAnotherExample \
	-only-testing:ManyTestsUITests/ManyTestsUITests/testAnotherExample \
	-only-testing:ManyTestsUITests/ManyTestsUITestsLaunchTests/testAnotherLaunch \
	-destination 'platform=iOS Simulator,name=iPhone 16 Pro Max,OS=latest' \
