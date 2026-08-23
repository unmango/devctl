# devctl

A CLI tool for development productivity.

## Installation

```shell
go install github.com/unmango/devctl/cmd
```

## Usage

The current supported functionalitly includes listing source code files, managing dependency version files, and driving stacked diffs.

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

### Stacked diffs

Thin orchestration over [gh-stack](https://gh.io/stacks), which must be installed:

```shell
gh extension install github/gh-stack
```

Each subcommand runs a fixed sequence of `git` and `gh stack` commands, nothing more.

| Command                                | Runs                                                                                                                       |
| -------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- |
| `devctl stack add <branch> -Am "msg"`  | `gh stack add`, or `gh stack init` plus a commit when there is no stack yet                                                |
| `devctl stack save -Am "msg"`          | `git add`, `git commit`, `gh stack rebase --upstack`, `gh stack push`                                                      |
| `devctl stack edit <branch> -Am "msg"` | `gh stack checkout`, `git add`, `git commit`, `gh stack rebase --upstack`, `gh stack checkout <original>`, `gh stack push` |
| `devctl stack ship`                    | `gh stack sync`, `gh stack submit --auto`                                                                                  |

```shell
$ devctl stack add feat/auth -Am "Add auth middleware"
$ devctl stack add feat/api -Am "Add API routes"
$ devctl stack save -Am "Handle expired tokens"
$ devctl stack edit feat/auth -Am "Fix the middleware"  # commits down the stack, returns to where you were
$ devctl stack ship --open
```

`-A`, `-u`, and `-m` mirror their `gh stack add` meanings, and staging or committing is skipped when none are given.
`gh stack` exit codes are passed through untouched, so a rebase conflict still exits 3.

Two things worth knowing:

- Pass `--remote` when the repo has more than one remote and `remote.pushDefault` isn't configured.
- `gh stack sync` succeeds while printing `Sync aborted` when the local and remote stacks have diverged, so `ship` cannot detect that case and will go on to submit.

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
