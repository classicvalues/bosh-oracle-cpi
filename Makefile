default: build

# Builds bosh-cpi for linux-amd64
build:
	go build -o out/cpi github.com/oracle/bosh-oracle-cpi/main

# Build cross-platform binaries
build-all:
	gox -output="out/cpi_{{.OS}}_{{.Arch}}" oracle/bosh-cpi/main

# Prepration for tests
get-deps:
	# Go lint tool
	go get github.com/golang/lint/golint

	# Simplify cross-compiling
	go get github.com/mitchellh/gox

	# Ginkgo and omega test tools
	go get github.com/onsi/ginkgo/ginkgo
	go get github.com/onsi/gomega

# Cleans up directory and source code with gofmt
clean:
	go clean ./...

# Run gofmt on all code
fmt:
	gofmt -l -w .

# Run linter with non-stric checking
lint:
	@echo ls -d */ | grep -v vendor | xargs -L 1 golint
	ls -d */ | grep -v vendor | xargs -L 1 golint

# Vet code
vet:
	go tool vet $$(ls -d */ | grep -v vendor)

# Runs the unit tests with coverage
test: get-deps clean fmt lint vet build
	ginkgo -r -race -skipPackage=bmc/test .

# Runs BMC integration tests
bmctest: get-deps
	# Uncomment and export variables (in the launch shell) to control test configuration
	# CPITEST_CONFIG=/path/to/my/oraclebmc/config ini. Default is ~/.oraclebmc/config
	# CPITEST_PROFILE=section inside CPITEST_CONFIG file. Default is CPITEST
	ginkgo bmc/test -slowSpecThreshold=500 -progress -nodes=3 -randomizeAllSpecs -randomizeSuites $(GINKGO_ARGS) -v

# Runs the integration tests from Concourse
testintci: get-deps bmctest


# Checks and creates, if necessary, resources in a project required to run integration tests.
configint:
	@echo "TODO: Integrate Terraform templates to setup test compartment and network"
	exit 1

# Deletes the resources created by the configint target
cleanint: check-proj
	@echo "TODO: Remove resources created by configint"
	exit 1