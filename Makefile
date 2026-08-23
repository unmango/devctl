NIX ?= nix

.PHONY: build check fmt test e2e tidy dev

build:
	$(NIX) build

check:
	$(NIX) flake check --all-systems

fmt:
	$(NIX) fmt

test:
	$(NIX) run .#test

e2e:
	$(NIX) run .#e2e

tidy:
	$(NIX) run .#tidy

dev:
	$(NIX) develop

# ginkgo scaffolding
%_suite_test.go:
	cd $(dir $@) && go tool ginkgo bootstrap

%_test.go:
	cd $(dir $@) && go tool ginkgo generate $(notdir $*)
