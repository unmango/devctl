# devctl

A CLI tool for development productivity.

## Installation

```shell
go install github.com/unmango/devctl/cmd
```

## Usage

The current supported functionalitly includes listing source code files and managing dependency version files.

### List source files

Useful for `make` targets

```shell
$ devctl list --go
# cmd/devops_suite_test.go
# cmd/init_test.go
# cmd/main.go
# cmd/version_test.go
# ...
```

```shell
$ devctl list --go --exclude-tests
# cmd/main.go
# pkg/cmd/init/version.go
# pkg/cmd/init.go
# ...
```

### Dependency versioning via files

This is a convention based versioning system where version numbers are stored in plaintext files located in the `.versions` directory of a given repository's root.

```shell
$ devctl init version foo v0.0.69 && cat .versions/foo
# 0.0.69
```

```shell
$ devctl version foo
# 0.0.69
```

This convention is useful alongside `make` where version updates can trigger targets:

```make
bin/mybin: .versions/mybin
    go install mybin@v$(shell cat $<)
```
