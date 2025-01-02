_ := $(shell mkdir -p .make bin)

WORKING_DIR := $(shell pwd)
LOCALBIN    := ${WORKING_DIR}/bin

export GOBIN := ${LOCALBIN}

GINKGO := ${LOCALBIN}/ginkgo

ifeq ($(shell test -f ${LOCALBIN}/devctl && echo yes),yes)
DEVCTL := ${LOCALBIN}/devctl
else
DEVCTL := go run ./cmd
endif

ifeq ($(CI),)
TEST_FLAGS := --label-filter !E2E
else
TEST_FLAGS := --github-output --race --trace --coverprofile=cover.profile
endif

build: bin/devctl
tidy: go.sum

test: .make/test
test_all:
	$(GINKGO) run -r ./

bin/devctl: $(shell $(DEVCTL) list --go --exclude-tests)
	go build -o $@ ./cmd

bin/ginkgo: go.mod
	go install github.com/onsi/ginkgo/v2/ginkgo

go.sum: go.mod $(shell $(DEVCTL) list --go)
	go mod tidy

%_suite_test.go: | bin/ginkgo
	cd $(dir $@) && $(GINKGO) bootstrap

%_test.go: | bin/ginkgo
	cd $(dir $@) && $(GINKGO) generate $(notdir $*)

.envrc: hack/example.envrc
	cp $< $@

.make/test: $(shell $(DEVCTL) list --go) | bin/ginkgo
	$(GINKGO) run ${TEST_FLAGS} $(sort $(dir $?))
	@touch $@
