# AGENTS.md

This file provides guidance to coding agents when working with code in this repository.

## Commands

Everything goes through Nix.
The `Makefile` targets are thin aliases, and `direnv` (`.envrc` runs `use flake`) puts the dev shell on PATH.

| Command           | Alias        | Does                                                          |
| ----------------- | ------------ | ------------------------------------------------------------- |
| `nix flake check` | `make check` | Unit tests, formatting, linting. The one command CI gates on. |
| `nix build`       | `make build` | Builds `result/bin/devctl`.                                   |
| `nix fmt`         | `make fmt`   | treefmt: gofmt, nixfmt, dprint, actionlint.                   |
| `nix run .#test`  | `make test`  | Unit tests only, unsandboxed, faster loop.                    |
| `nix run .#e2e`   | `make e2e`   | E2E tests. Needs network.                                     |
| `nix run .#tidy`  | `make tidy`  | `go mod tidy` + `gomod2nix` after any `go.mod` change.        |

`make` is shadowed by a Prezto autoload stub in this shell, so prefix it: `command make check`.

### Running a subset of tests

`nix run .#test` and `nix run .#e2e` forward trailing args to `ginkgo run -r`, so a path narrows the run:

```shell
nix run .#test -- ./pkg/stack        # one package
go tool ginkgo run --label-filter='!E2E' --focus='Add' ./pkg/stack
```

Inside the dev shell, use `go tool ginkgo` rather than a `ginkgo` on PATH.
It builds the version pinned in `go.mod` and avoids a version skew warning.

New ginkgo files are scaffolded by pattern rules in the `Makefile`:

```shell
command make pkg/foo/foo_suite_test.go   # ginkgo bootstrap
command make pkg/foo/bar_test.go         # ginkgo generate bar
```

### The E2E label gotcha

Unit and E2E runs are split by ginkgo label, not by directory.
Only specs explicitly decorated `Label("E2E")` (currently just `test/e2e/install_test.go`) are excluded from the unit run.
Every other spec under `test/e2e/` runs during `nix run .#test` and inside the sandboxed `checks.unit` derivation.
A new spec in `test/e2e/` that needs the network must carry `Label("E2E")` or it will break `nix flake check`.

## Architecture

### cmd/ is cobra wiring, pkg/ is the logic

`cmd/` packages only parse flags, resolve a working directory, call into `pkg/`, and exit.
No behavior lives there beyond flag definitions and error-to-exit-code mapping.
Every command is built by a `NewX()` constructor with a package level `var XCmd = NewX()`, so tests can build a fresh command with its own options struct.

`main.go` does one thing before cobra sees the args: if invoked with exactly one argument that looks like a `.versions/...` path, it prints that version and returns.
This is what makes `$(shell devctl $<)` work in a Makefile prerequisite.

### work.Directory is the filesystem boundary

`pkg/work.Directory` is a string path with an `afero.Fs` rooted at it (`afero.NewBasePathFs`).
`work.Load` resolves the git root first, then falls back to cwd.
`work.ChdirOptions` / `work.ChdirFlag` add the shared `-C`/`--chdir` flag; commands call `opts.Chdir(ctx)` before doing anything.

Filesystem access goes through `afero` so tests can substitute a mem-map fs.
`pkg/config` is the exception, its viper instance still uses the real fs (there is a `TODO` about it).

### Functional options, split across three packages

`pkg/version` uses a three-package option layout worth copying if you extend it:

- `pkg/version/opts` holds the public option structs, defaults, and `XOp` functions.
- `pkg/version/internal` applies them via `unmango/go/fopt.ApplyAll` and returns the resolved struct.
- `pkg/version` takes `...opts.XOp` variadics.

Simpler packages (`pkg/stack`, `pkg/agents`) just take a plain options struct by value.

### pkg/stack shells out and gets out of the way

`pkg/stack` never reimplements `gh stack`.
Each exported function is a fixed sequence of `git` and `gh stack` invocations with stdin/stdout/stderr wired straight through, so the underlying tools own all output, prompting, and exit codes.
`cmd/stack.fail` propagates the child's exit code (2 not in a stack, 3 rebase conflict, 7 rebase in progress) instead of collapsing it to 1.

Binary resolution is layered: context value (`WithGhPath`/`WithGitPath`), then `GH_PATH`/`GIT_PATH` env vars, then `exec.LookPath`.
Tests use the context layer.

### Testing style

Ginkgo v2 + Gomega throughout, one `*_suite_test.go` per package, external test packages (`package stack_test`).

`pkg/stack` tests use a stub harness in `stack_suite_test.go`: `NewStub` writes a shell script that logs each invocation pipe-separated and replies from a `map[string]Response` keyed on the exact argument string.
Assertions compare `gh.Calls()` against the literal expected command sequence.
This is the pattern for anything that shells out, prefer it over mocking `exec`.

`test/e2e` builds the real binary with `gexec.Build` in `BeforeSuite` and drives it against a temp dir populated from an embedded `testdata` FS.

## Releasing

release-please cuts releases from Conventional Commits.
PRs are squash-merged, so the **PR title** is the commit subject release-please reads, and `.github/workflows/pr-title.yml` rejects non-conventional titles.

`feat:` bumps minor, `fix:` bumps patch.
The version lives in `.release-please-manifest.json` and `nix/packages.nix` imports it; never bump a version by hand.

## Gotchas

- Any `go.mod` edit must be followed by `nix run .#tidy`, CI fails on a dirty `gomod2nix.toml`.
- CI also runs `git diff --exit-code` on the whole tree, so nothing may be generated at build time and left uncommitted.
- Flake outputs live in `nix/*.nix`; `flake.nix` only declares inputs.
- `checks.unit` overrides the `devctl` package with `doCheck = true`, so package build stays fast and a test failure is attributable to the check.
