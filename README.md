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

### Agent instructions

`claude /init` writes a `CLAUDE.md`, but the build commands and conventions in it are useful to every coding agent.
`devctl migrate agents` promotes that file to the portable [`AGENTS.md`](https://agents.md) convention and leaves pointers behind for the tools looking for their own filename.

```shell
$ devctl migrate agents
# AGENTS.md                        the instructions, boilerplate rewritten
# CLAUDE.md                        imports AGENTS.md
# .github/copilot-instructions.md  points at AGENTS.md
```

Only the fixed boilerplate is rewritten: the title, the `Claude Code (claude.ai/code)` guidance sentence, and `CLAUDE.md` self references.
Everything else is left alone, and any line still mentioning Claude is reported so it can be edited by hand.

```shell
$ devctl migrate agents
# AGENTS.md:9 still mentions Claude: Ask Claude to run the tests.
```

Existing files are never clobbered without `--force`, and a `CLAUDE.md` already importing `AGENTS.md` is refused outright so a second run cannot promote the pointer over the instructions it points at.

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

### Releasing

Releases are cut by [release-please](https://github.com/googleapis/release-please) from [Conventional Commits](https://www.conventionalcommits.org).
PRs are squash-merged, so the **PR title** is the commit subject release-please reads; a CI check rejects titles that aren't conventional.

`feat:` bumps the minor version, `fix:` the patch; anything else lands in the changelog without forcing a release.
release-please keeps an open release PR with the pending `CHANGELOG.md` and version bump.
Merging it tags `vX.Y.Z`, publishes the GitHub Release, and triggers goreleaser to attach the binaries.

The released version is read from `.release-please-manifest.json`; nothing needs to be bumped by hand.
