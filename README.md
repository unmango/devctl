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
# cmd/devctl_suite_test.go
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
    go install mybin@$(shell devctl $<)
```

## Development

This repo is driven by Nix.
[direnv](https://direnv.net) picks up the checked-in `.envrc` and drops you into the dev shell; without it, run `nix develop`.

| Command           | Equivalent   | Does                                                                |
| ----------------- | ------------ | ------------------------------------------------------------------- |
| `nix flake check` | `make check` | Unit tests, formatting, and linting. The one command CI runs.       |
| `nix build`       | `make build` | Builds the `devctl` binary to `result/bin/devctl`.                  |
| `nix fmt`         | `make fmt`   | Formats everything via treefmt (gofmt, nixfmt, dprint, actionlint). |
| `nix run .#test`  | `make test`  | Unit tests only, outside the sandbox for a faster loop.             |
| `nix run .#e2e`   | `make e2e`   | End-to-end tests. Needs network access.                             |
| `nix run .#tidy`  | `make tidy`  | Syncs `go.sum` and `gomod2nix.toml` after a `go.mod` change.        |

Flake outputs live in [`nix/`](./nix); `flake.nix` only declares inputs.
The released version is read from `.versions/devctl`.
