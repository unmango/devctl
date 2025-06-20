_ := $(shell mkdir -p .make bin)

WORKING_DIR := $(shell pwd)
LOCALBIN    := ${WORKING_DIR}/bin

export GOBIN := ${LOCALBIN}

GINKGO  := go tool ginkgo
JSON2GO := ${LOCALBIN}/go-jsonschema
JQ      := ${LOCALBIN}/jq

ifeq ($(shell test -f ${LOCALBIN}/devctl && echo yes),yes)
DEVCTL := ${LOCALBIN}/devctl
else
DEVCTL := go run ./
endif

ifeq ($(CI),)
TEST_FLAGS := --label-filter !E2E
else
TEST_FLAGS := --github-output --race --trace --coverprofile=cover.profile
endif

-load bin/mk_funcs.so

build: bin/devctl bin/mk_funcs.so
tidy: go.sum
format: .make/dprint-format

test: .make/test
test_all:
	$(GINKGO) run -r ./

bin/devctl: $(shell $(DEVCTL) list --go --exclude-tests)
	go build -o $@ ./

bin/mk_funcs.so:
	go build -o $@ -buildmode=c-shared ./pkg/make/mk_funcs

bin/go-jsonschema: .versions/go-jsonschema
	go install github.com/atombender/go-jsonschema@$(shell $(DEVCTL) $<)

bin/jq: .versions/jq
	curl -L -o $@ https://github.com/jqlang/jq/releases/download/jq-$(shell $(DEVCTL) v jq)/jq-$(shell go env GOOS | sed s/darwin/macos/)-$(shell go env GOARCH)
	chmod +x $@

go.sum: go.mod $(shell $(DEVCTL) list --go)
	go mod tidy

# I can't seem to get --schema-root-type to do what I want it to
pkg/renovate/zz_generated.schema.go: .make/renovate-schema.json bin/go-jsonschema
	mkdir -p $(dir $@)
	$(JSON2GO) --package renovate $< --only-models | sed s/RenovateSchemaJson/Config/g > $@

%_suite_test.go:
	cd $(dir $@) && $(GINKGO) bootstrap

%_test.go:
	cd $(dir $@) && $(GINKGO) generate $(notdir $*)

.envrc: hack/example.envrc
	cp $< $@

.make/test: $(shell $(DEVCTL) list --go)
	$(GINKGO) run ${TEST_FLAGS} $(sort $(dir $?))
	@touch $@

.make/renovate-schema.orig.json:
	curl https://docs.renovatebot.com/renovate-schema.json -o $@

.make/renovate-schema.json: .make/renovate-schema.orig.json hack/renovate/*.jq | bin/jq
	cat $< | $(JQ) -f hack/renovate/delete-refs.jq > $@

.make/dprint-format: .dprint.json README.md .github/renovate.json .vscode/extensions.json
	dprint fmt
